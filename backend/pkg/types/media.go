package types

type Media string

func (m Media) String() string  { return string(m) }
func (m Media) Kind() FieldType { return FieldMedia }

func (m Media) IsImage() bool { return Image(m).IsImage() }
func (m Media) IsVideo() bool { return Video(m).IsVideo() }
func (m Media) IsAudio() bool { return Audio(m).IsAudio() }
func (m Media) IsMedia() bool { return Media(m).IsMedia() }

func (m Media) Ext() string {
	return Image(m).Ext() // same logic
}

func (m Media) Type() FieldType {
	if m.IsImage() {
		return FieldImage
	}
	if m.IsVideo() {
		return FieldVideo
	}
	if m.IsAudio() {
		return FieldAudio
	}
	return FieldFile
}
