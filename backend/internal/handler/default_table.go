package handler

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/types"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func GET_DEFAULT_TableDataHandler(db *gorm.DB, model interface{}, preload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := reflect.TypeOf(model).Elem()
		// =============================
		// 🔹 Handle Request All Data
		// =============================
		if len(c.Request.URL.Query()) == 0 {
			// fmt.Println("No query parameters in URL")
			// fallback to query all if query binding failed
			results := reflect.New(reflect.SliceOf(t)).Interface()
			if len(preload) > 0 {
				for _, p := range preload {
					db = db.Preload(p)
				}
				db = db.Find(results)
			} else {
				db = db.Find(results)
			}
			sliceValue := reflect.ValueOf(results).Elem()
			data := make([]gin.H, 0, sliceValue.Len())

			for i := 0; i < sliceValue.Len(); i++ {
				row := sliceValue.Index(i)
				rowData := gin.H{}
				for j := 0; j < t.NumField(); j++ {
					field := t.Field(j)
					if jsonKey := field.Tag.Get("json"); jsonKey != "" && jsonKey != "-" {
						if row.Field(j).Type() == reflect.TypeOf(time.Time{}) {
							rowData[jsonKey] = row.Field(j).Interface().(time.Time).Format(util.T_YYYYMMDD_HHmmss)
						} else if row.Field(j).Type() == reflect.TypeOf(sql.NullTime{}) {
							if row.Field(j).Interface().(sql.NullTime).Valid {
								rowData[jsonKey] = row.Field(j).Interface().(sql.NullTime).Time.Format(util.T_YYYYMMDD_HHmmss)
							} else {
								rowData[jsonKey] = ""
							}
						} else {
							rowData[jsonKey] = row.Field(j).Interface()
						}
					}

				}
				data = append(data, rowData)
			}
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Fetched all records (fallback, no query found)",
				"data":    data,
			})
			return
		}
		// =============================
		// 🔹 Handle Download Batch Upload Template
		// =============================
		if _, exists := c.GetQuery("batch_upload_template.xlsx"); exists {

			// Create a new Excel file in memory
			f := excelize.NewFile()
			sheetName := "Sheet1"

			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			structName := t.Name()
			f.SetCellValue(sheetName, "A1", "Batch Upload "+util.AddSpaceBeforeUppercase(structName))
			// var tableHeaders []string
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				jsonKey := field.Tag.Get("json")
				if jsonKey == "" || jsonKey == "-" {
					continue
				}
				varName := field.Name
				switch varName {
				case "Id", "ID":
					continue
				}
				fillInfo := ""
				fieldType := field.Type
				if fieldType == reflect.TypeOf(time.Time{}) {
					time_format := field.Tag.Get("time_format")
					if time_format != "" {
						humanReadableFormat := strings.ReplaceAll(time_format, "20", "YY")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "06", "YY")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "15", "HH")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "04", "mm")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "05", "ss")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "01", "MM")
						humanReadableFormat = strings.ReplaceAll(humanReadableFormat, "02", "DD")
						fillInfo = "(" + humanReadableFormat + ")"
					} else {
						fillInfo = "(YYYY-MM-DD)(YYYY-MM-DD HH:mm)"
					}
				}
				// Add data to specific cells
				f.SetCellValue(sheetName, util.NumberToAlphabet(i)+"2", util.AddSpaceBeforeUppercase(field.Name)+" "+fillInfo)
			}

			// Write the file content to an in-memory buffer
			var buffer bytes.Buffer
			if err := f.Write(&buffer); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": fmt.Sprintf("Failed to create Excel file: %v", err),
					"error":   fmt.Sprintf("Failed to create Excel file: %v", err),
				})
				return
			}

			// Set the necessary headers for file download
			c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=batch_upload_%s.xlsx", util.ToSnakeCase(structName)))

			// Stream the Excel file to the response
			_, err := c.Writer.Write(buffer.Bytes())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": fmt.Sprintf("Failed to write Excel file to response: %v", err),
					"error":   fmt.Sprintf("Failed to write Excel file to response: %v", err),
				})
			}
			return
		}
		// =============================
		// 🔹 Handle Filtering Request
		// =============================
		var request struct {
			Draw       int    `form:"draw"`
			Start      int    `form:"start"`
			Length     int    `form:"length"`
			Search     string `form:"search[value]"`
			SortColumn int    `form:"order[0][column]"`
			SortDir    string `form:"order[0][dir]"`
		}

		if err := c.BindQuery(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		// =============================
		// 🔹 Siapkan query dasar
		// =============================
		filteredQuery := db.Model(model)
		if len(preload) > 0 {
			for _, p := range preload {
				filteredQuery = filteredQuery.Preload(p)
			}
		}
		// =============================
		// 🔹 Ambil kolom yang dikirim dari frontend
		// =============================
		columnMap := make(map[int]string)
		columnSearch := make(map[int]string)

		jsonTagList := []string{}
		gormTagList := []string{}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			jsonTag := field.Tag.Get("json")
			gormTag := field.Tag.Get("gorm")
			if jsonTag == "" || jsonTag == "-" {
				continue
			}
			jsonTagList = append(jsonTagList, jsonTag)
			columnKey := ""
			for tag := range strings.SplitSeq(gormTag, ";") {
				if strings.HasPrefix(tag, "column:") {
					columnKey = strings.TrimPrefix(tag, "column:")
					break
				}
			}
			if columnKey == "" {
				columnKey = jsonTag
			}
			gormTagList = append(gormTagList, columnKey)
		}
		for i := 0; i < t.NumField(); i++ {
			colKey := fmt.Sprintf("columns[%d][data]", i)
			colValue := c.Query(colKey)
			if colValue == "" {
				continue // DataTables kirim kolom berurutan dari 0
			}
			if !util.Contains(jsonTagList, colValue) {
				continue
			}
			columnMap[i] = colValue

			searchVal := c.Query(fmt.Sprintf("columns[%d][search][value]", i))
			if searchVal != "" {
				columnSearch[i] = searchVal
			}
		}

		// =============================
		// 🔹 Tentukan kolom sorting
		// =============================
		sortColumnName := columnMap[request.SortColumn]
		if sortColumnName == "" {
			sortColumnName = "id" // fallback
		}
		orderString := fmt.Sprintf("`%s` %s", sortColumnName, request.SortDir)

		// =============================
		// 🔹 Prioritaskan search spesifik kolom
		// =============================
		hasColumnSearch := len(columnSearch) > 0
		if hasColumnSearch {
			for idx, val := range columnSearch {
				if val == "" {
					continue
				}
				columnName := columnMap[idx]
				if columnName == "" {
					continue
				}
				if !util.Contains(gormTagList, columnName) {
					continue
				}
				filteredQuery = filteredQuery.Where(fmt.Sprintf("%s LIKE ?", columnName), "%"+val+"%")
			}
		} else if request.Search != "" {
			// =============================
			// 🔹 Kalau tidak ada search per kolom → gunakan search global
			// =============================
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)

				jsonTag := field.Tag.Get("json")
				gormTag := field.Tag.Get("gorm")

				columnKey := ""
				for _, tag := range strings.Split(gormTag, ";") {
					if strings.HasPrefix(tag, "column:") {
						columnKey = strings.TrimPrefix(tag, "column:")
						break
					}
				}
				if columnKey == "" {
					columnKey = jsonTag
				}
				if columnKey == "" || columnKey == "-" {
					continue
				}

				filteredQuery = filteredQuery.Or(fmt.Sprintf("`%s` LIKE ?", columnKey), "%"+request.Search+"%")
			}
		}

		// =============================
		// 🔹 Hitung total dan filtered
		// =============================
		var totalRecords int64
		db.Model(model).Count(&totalRecords)

		var filteredRecords int64
		filteredQuery.Count(&filteredRecords)

		// =============================
		// 🔹 Apply sort + pagination
		// =============================
		query := filteredQuery.Order(orderString).Offset(request.Start).Limit(request.Length)

		results := reflect.New(reflect.SliceOf(reflect.TypeOf(model).Elem())).Interface()
		if err := query.Find(results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// =============================
		// 🔹 Format hasil JSON
		// =============================
		sliceValue := reflect.ValueOf(results).Elem()
		data := make([]gin.H, 0, sliceValue.Len())

		for i := 0; i < sliceValue.Len(); i++ {
			row := sliceValue.Index(i)
			rowData := gin.H{}
			for j := 0; j < t.NumField(); j++ {
				field := t.Field(j)
				if jsonKey := field.Tag.Get("json"); jsonKey != "" && jsonKey != "-" {
					if row.Field(j).Type() == reflect.TypeOf(time.Time{}) {
						rowData[jsonKey] = row.Field(j).Interface().(time.Time).Format(time.RFC3339) // Format ISO 8601
					} else if row.Field(j).Type() == reflect.TypeOf(sql.NullTime{}) {
						if row.Field(j).Interface().(sql.NullTime).Valid {
							rowData[jsonKey] = row.Field(j).Interface().(sql.NullTime).Time.Format(time.RFC3339) // Format ISO 8601
						} else {
							rowData[jsonKey] = ""
						}
					} else {
						rowData[jsonKey] = row.Field(j).Interface()
					}
				}

			}
			data = append(data, rowData)
		}

		// =============================
		// 🔹 Kirim ke frontend
		// =============================
		c.JSON(http.StatusOK, gin.H{
			"success":         true,
			"message":         "OK",
			"draw":            request.Draw,
			"recordsTotal":    totalRecords,
			"recordsFiltered": filteredRecords,
			"data":            data,
		})
	}
}

func DELETE_DEFAULT_TableDataHandler(db *gorm.DB, model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Req struct {
			ID []int `form:"id"`
		}
		var req Req
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(req.ID) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No ID provided", "error": "No ID provided"})
			return
		}
		if err := db.Delete(model, req.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func POST_DEFAULT_TableDataHandler(db *gorm.DB, model interface{}, preload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the model type
		instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

		// Bind JSON request body to the instance
		if err := c.ShouldBindJSON(instance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Create the record in database
		if err := db.Create(instance).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Optionally reload with preloaded relationships
		if len(preload) > 0 {
			query := db
			for _, p := range preload {
				query = query.Preload(p)
			}
			if err := query.First(instance, instance).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Record created but failed to reload: " + err.Error(),
				})
				return
			}
		}

		// Format the response
		t := reflect.TypeOf(model).Elem()
		v := reflect.ValueOf(instance).Elem()
		rowData := gin.H{}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if jsonKey := field.Tag.Get("json"); jsonKey != "" && jsonKey != "-" {
				fieldValue := v.Field(i)
				if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
					rowData[jsonKey] = fieldValue.Interface().(time.Time).Format(util.T_YYYYMMDD_HHmmss)
				} else if fieldValue.Type() == reflect.TypeOf(sql.NullTime{}) {
					if fieldValue.Interface().(sql.NullTime).Valid {
						rowData[jsonKey] = fieldValue.Interface().(sql.NullTime).Time.Format(util.T_YYYYMMDD_HHmmss)
					} else {
						rowData[jsonKey] = ""
					}
				} else {
					rowData[jsonKey] = fieldValue.Interface()
				}
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Record created successfully",
			"data":    rowData,
		})
	}
}
func PUT_DEFAULT_TableDataHandler(db *gorm.DB, model interface{}, preload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the model type
		instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

		// Bind JSON request body to the instance
		if err := c.ShouldBindJSON(instance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Create the record in database
		if err := db.Save(instance).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Optionally reload with preloaded relationships
		if len(preload) > 0 {
			query := db
			for _, p := range preload {
				query = query.Preload(p)
			}
			if err := query.First(instance, instance).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Record created but failed to reload: " + err.Error(),
				})
				return
			}
		}

		// Format the response
		t := reflect.TypeOf(model).Elem()
		v := reflect.ValueOf(instance).Elem()
		rowData := gin.H{}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if jsonKey := field.Tag.Get("json"); jsonKey != "" && jsonKey != "-" {
				fieldValue := v.Field(i)
				if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
					rowData[jsonKey] = fieldValue.Interface().(time.Time).Format(util.T_YYYYMMDD_HHmmss)
				} else if fieldValue.Type() == reflect.TypeOf(sql.NullTime{}) {
					if fieldValue.Interface().(sql.NullTime).Valid {
						rowData[jsonKey] = fieldValue.Interface().(sql.NullTime).Time.Format(util.T_YYYYMMDD_HHmmss)
					} else {
						rowData[jsonKey] = ""
					}
				} else {
					rowData[jsonKey] = fieldValue.Interface()
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Record Updated successfully",
			"data":    rowData,
		})
	}
}
func PATCH_DEFAULT_TableDataHandler(db *gorm.DB, model interface{}, preload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := reflect.TypeOf(model).Elem()

		// Extract ID from multiple sources (query, form, json)
		var ids []string

		// Try to get ID from query parameter
		if idQuery := c.QueryArray("id"); len(idQuery) > 0 {
			ids = append(ids, idQuery...)
		} else if idQuery := c.QueryArray("Id"); len(idQuery) > 0 {
			ids = append(ids, idQuery...)
		} else if idQuery := c.QueryArray("ID"); len(idQuery) > 0 {
			ids = append(ids, idQuery...)
		}

		// Parse raw request body into map
		var requestData map[string]interface{}
		contentType := c.ContentType()
		switch contentType {
		case "application/json":
			fmt.Println("JSON")
			if err := c.ShouldBindJSON(&requestData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid JSON: " + err.Error(),
				})
				return
			}
		case "application/x-www-form-urlencoded":
			fmt.Println("Form URL Encoded")
			if err := c.ShouldBind(&requestData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid JSON: " + err.Error(),
				})
				return
			}
		case "multipart/form-data":
			fmt.Println("Multipart Form")
			if len(ids) == 0 && (c.PostFormArray("id") == nil && c.PostFormArray("Id") == nil && c.PostFormArray("ID") == nil) {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "ID is required in query or form data",
				})
				return
			}

		default:
			fmt.Println("Unknown content type:", contentType)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Unsupported content type: " + contentType,
			})
			return
		}

		// Try to get ID from request data (JSON/form)
		if idVal, ok := requestData["id"]; ok && idVal != nil {
			ids = append(ids, fmt.Sprintf("%v", idVal))
		} else if idVal, ok := requestData["Id"]; ok && idVal != nil {
			ids = append(ids, fmt.Sprintf("%v", idVal))
		} else if idVal, ok := requestData["ID"]; ok && idVal != nil {
			ids = append(ids, fmt.Sprintf("%v", idVal))
		}

		// Check if we have at least one ID
		if len(ids) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "ID is required (id/Id/ID in query, form, or JSON)",
			})
			return
		}
		ids = util.Unique(ids)

		// Build update map by matching JSON keys with model fields
		updateMap := make(map[string]interface{})
		validate := validator.New()
		errorExist := false
		errorList := []string{}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			gormTag := field.Tag.Get("gorm")

			if jsonTag == "" || jsonTag == "-" || strings.ToLower(jsonTag) == "id" {
				continue
			}
			if ui := types.ParseUIOptions(field.Tag.Get("ui")); !ui.Editable {
				continue
			}
			// Get the column name from GORM tag or use json tag
			columnName := jsonTag
			if gormTag != "" {
				for _, tag := range strings.Split(gormTag, ";") {
					if strings.HasPrefix(tag, "column:") {
						columnName = strings.TrimPrefix(tag, "column:")
						break
					}
				}
			}

			if contentType == "multipart/form-data" {
				dataType := string(types.DetectFieldType(field.Type))
				switch dataType {
				case "string", "text", "number", "password", "badge":
					if val := c.PostForm(jsonTag); val != "" {
						requestData[jsonTag] = val
					}
				case "email":
					if val := c.PostForm(jsonTag); val != "" {
						if validate.Var(val, "required,email") == nil {
							requestData[jsonTag] = val
						} else {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid email format for field %s", jsonTag))
						}
					}
				case "phone":
					if val := c.PostForm(jsonTag); val != "" {
						// contoh rule: hanya angka, panjang min 8
						if err := validate.Var(val, "numeric,min=8"); err == nil {
							requestData[jsonTag] = val
						} else {
							errorExist = true
							errorList = append(errorList,
								fmt.Sprintf("Invalid phone number for field %s (only numbers, min 8 digits)", jsonTag),
							)
						}
					}
				case "datetime", "date", "time":
					formats := map[string]struct {
						layout string
						errMsg string
					}{
						"datetime": {"2006-01-02 15:04:05", "Invalid datetime format for field %s (use YYYY-MM-DD HH:MM:SS)"},
						"date":     {"2006-01-02", "Invalid date format for field %s (use YYYY-MM-DD)"},
						"time":     {"15:04:05", "Invalid time format for field %s (use HH:MM:SS)"},
					}
					cfg := formats[dataType]
					val := c.PostForm(jsonTag)

					if val != "" {
						if _, err := time.Parse(cfg.layout, val); err == nil {
							requestData[jsonTag] = val
						} else {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf(cfg.errMsg, jsonTag))
						}
					}
				case "boolean":
					if val, ok := c.GetPostForm(jsonTag); ok {
						boolVal := true
						if val == "false" || val == "0" || val == "" {
							boolVal = false
						}
						requestData[jsonTag] = boolVal
					}
				case "avatar", "image", "file", "video", "audio", "document", "media", "archive":
					file, err := c.FormFile(jsonTag)
					if err != nil {
						continue
					}
					switch dataType {
					case "avatar", "image":
						if !types.Image(file.Filename).IsImage() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					case "video":
						if !types.Video(file.Filename).IsVideo() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					case "audio":
						if !types.Audio(file.Filename).IsAudio() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					case "document":
						if !types.Document(file.Filename).IsDocument() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					case "media":
						if !types.Media(file.Filename).IsMedia() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					case "archive":
						if !types.Archive(file.Filename).IsArchive() {
							errorExist = true
							errorList = append(errorList, fmt.Sprintf("Invalid %s file for field %s", dataType, jsonTag))
						}
					}
					if errorExist {
						continue
					}
					fileStorePath, err := util.GetAppDataDir(os.Getenv("APP_NAME"))
					if err != nil {
						fileStorePath = "./uploads"
					}
					savedFilePath := filepath.Join(fileStorePath, strings.ToLower(t.Name()))
					if err := os.MkdirAll(savedFilePath, os.ModePerm); err != nil {
						errorExist = true
						errorList = append(errorList,
							fmt.Sprintf("Internal Error for field %s, %s", jsonTag, err.Error()),
						)
						continue
					}
					fullPath := filepath.Join(savedFilePath, fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename)))
					if err := c.SaveUploadedFile(file, fullPath); err != nil {
						errorExist = true
						errorList = append(errorList,
							fmt.Sprintf("Internal Error for field %s, %s", jsonTag, err.Error()),
						)
						continue
					}
					requestData[jsonTag] = fullPath
				}
			}
			// Check if this field exists in request data
			if val, exists := requestData[jsonTag]; exists {
				updateMap[columnName] = val
			}
		}
		if errorExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   strings.Join(errorList, "; "),
			})
			return
		}

		// If no fields to update
		if len(updateMap) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "No valid fields to update",
			})
			return
		}

		// Perform update using GORM Updates with WHERE IN
		result := db.Model(model).Where("id IN ?", ids).Updates(updateMap)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Update failed: " + result.Error.Error(),
			})
			return
		}

		// Check if any rows were affected
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "No records found with the provided ID(s)",
			})
			return
		}

		// Fetch updated records
		results := reflect.New(reflect.SliceOf(t)).Interface()
		query := db.Model(model)

		if len(preload) > 0 {
			for _, p := range preload {
				query = query.Preload(p)
			}
		}

		if err := query.Where("id IN ?", ids).Find(results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Records updated but failed to reload: " + err.Error(),
			})
			return
		}

		// Format the response
		sliceValue := reflect.ValueOf(results).Elem()
		data := make([]gin.H, 0, sliceValue.Len())

		for i := 0; i < sliceValue.Len(); i++ {
			row := sliceValue.Index(i)
			rowData := gin.H{}

			for j := 0; j < t.NumField(); j++ {
				field := t.Field(j)
				if jsonKey := field.Tag.Get("json"); jsonKey != "" && jsonKey != "-" {
					fieldValue := row.Field(j)
					if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
						rowData[jsonKey] = fieldValue.Interface().(time.Time).Format(util.T_YYYYMMDD_HHmmss)
					} else if fieldValue.Type() == reflect.TypeOf(sql.NullTime{}) {
						if fieldValue.Interface().(sql.NullTime).Valid {
							rowData[jsonKey] = fieldValue.Interface().(sql.NullTime).Time.Format(util.T_YYYYMMDD_HHmmss)
						} else {
							rowData[jsonKey] = ""
						}
					} else {
						rowData[jsonKey] = fieldValue.Interface()
					}
				}
			}
			data = append(data, rowData)
		}

		message := fmt.Sprintf("%d record(s) updated successfully", result.RowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"success":       true,
			"message":       message,
			"rows_affected": result.RowsAffected,
			"data":          data,
		})
	}
}
