# Configurable Table Usage

The `ConfigurableTable` component is a fully dynamic, config-driven data table that reads its configuration from localStorage and auto-generates filters, search inputs, and columns based on the table settings.

## Features

1. **Auto-generated Columns**: Reads column configuration from localStorage and generates TanStack Table columns
2. **Dynamic Filters**: 
   - Columns with `filterable: true` and `selection` URL automatically fetch filter options from backend
   - Columns with `filterable: true` and empty `selection` support text search
3. **Smart Search**:
   - Single text-searchable column: Uses column-specific search
   - Multiple text-searchable columns: Uses global search across all columns
4. **Dynamic Icons**: Filter options can include Lucide icon names that are dynamically loaded
5. **Client/Server Mode**: Supports both client-side and server-side data fetching

## Table Configuration Format

Configuration is stored in localStorage under key `app_table`:

```json
{
  "users": {
    "column": [
      {
        "name": "Email",
        "data": "email",
        "type": "email",
        "visible": true,
        "visibility": true,
        "sortable": true,
        "filterable": true,
        "editable": true,
        "passwordable": false,
        "selection": ""  // Empty = text search
      },
      {
        "name": "Status",
        "data": "status",
        "type": "badge",
        "visible": true,
        "visibility": true,
        "sortable": true,
        "filterable": true,
        "editable": true,
        "passwordable": false,
        "selection": "/options?data=status"  // URL = dropdown filter
      }
    ],
    "row": 10,
    "row_opt": [10, 20, 30, 40, 50],
    "sort": "id desc",
    "table_name": "users",
    "url": "/users"
  }
}
```

## Filter Options API

For columns with a `selection` URL, the component fetches filter options:

**Request:**
```
GET /options?data=status
```

**Response:**
```json
{
  "data": "status",
  "message": "",
  "options": [
    {
      "label": "Active",
      "value": "active",
      "icon": "check-circle"
    },
    {
      "label": "Inactive",
      "value": "inactive",
      "icon": "x-circle"
    }
  ],
  "success": true,
  "title": "Status"
}
```

## Usage Example

```tsx
import { ConfigurableTable } from '@/components/ConfigurableTable'

export function UsersPage() {
  return (
    <ConfigurableTable
      tableName="users"
      mode="server"  // or "client"
    />
  )
}
```

## Column Types

Supported column types:
- `text`: Plain text
- `email`: Email link
- `datetime`: Formatted date/time
- `number`: Numeric value
- `image`: Image thumbnail (h-10 w-10)
- `avatar`: Circular avatar (size-8)
- `badge`: Colored badge with auto-color mapping

## Filter Behavior

| Column Config | Behavior |
|---------------|----------|
| `filterable: true` + `selection: "/options?data=role"` | Dropdown filter with options from API |
| `filterable: true` + `selection: ""` | Text search (column-specific or global) |
| `filterable: false` | No filter |

## Search Behavior

- **1 text-searchable column**: Column-specific search input
- **Multiple text-searchable columns**: Global search across all columns
- **0 text-searchable columns**: Global search (default behavior)

## Icon Support

Filter options can include Lucide icon names:
- Icons are dynamically loaded using `lucide-react/dynamic`
- Valid icon names: https://lucide.dev/icons
- Example: `"sparkle"`, `"check-circle"`, `"user"`, etc.
