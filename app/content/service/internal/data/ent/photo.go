// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-blog/app/content/service/internal/data/ent/photo"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Photo is the model entity for the Photo schema.
type Photo struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *int64 `json:"delete_time,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Thumbnail holds the value of the "thumbnail" field.
	Thumbnail *string `json:"thumbnail,omitempty"`
	// 更新时间
	TakeTime *int64 `json:"take_time,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// Team holds the value of the "team" field.
	Team *string `json:"team,omitempty"`
	// Location holds the value of the "location" field.
	Location *string `json:"location,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Likes holds the value of the "likes" field.
	Likes *int32 `json:"likes,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Photo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case photo.FieldID, photo.FieldCreateTime, photo.FieldUpdateTime, photo.FieldDeleteTime, photo.FieldTakeTime, photo.FieldLikes:
			values[i] = new(sql.NullInt64)
		case photo.FieldName, photo.FieldThumbnail, photo.FieldURL, photo.FieldTeam, photo.FieldLocation, photo.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Photo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Photo fields.
func (ph *Photo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case photo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = uint32(value.Int64)
		case photo.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ph.CreateTime = new(int64)
				*ph.CreateTime = value.Int64
			}
		case photo.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ph.UpdateTime = new(int64)
				*ph.UpdateTime = value.Int64
			}
		case photo.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				ph.DeleteTime = new(int64)
				*ph.DeleteTime = value.Int64
			}
		case photo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ph.Name = new(string)
				*ph.Name = value.String
			}
		case photo.FieldThumbnail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail", values[i])
			} else if value.Valid {
				ph.Thumbnail = new(string)
				*ph.Thumbnail = value.String
			}
		case photo.FieldTakeTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field take_time", values[i])
			} else if value.Valid {
				ph.TakeTime = new(int64)
				*ph.TakeTime = value.Int64
			}
		case photo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				ph.URL = new(string)
				*ph.URL = value.String
			}
		case photo.FieldTeam:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team", values[i])
			} else if value.Valid {
				ph.Team = new(string)
				*ph.Team = value.String
			}
		case photo.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				ph.Location = new(string)
				*ph.Location = value.String
			}
		case photo.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ph.Description = new(string)
				*ph.Description = value.String
			}
		case photo.FieldLikes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likes", values[i])
			} else if value.Valid {
				ph.Likes = new(int32)
				*ph.Likes = int32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Photo.
// Note that you need to call Photo.Unwrap() before calling this method if this Photo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Photo) Update() *PhotoUpdateOne {
	return (&PhotoClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the Photo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Photo) Unwrap() *Photo {
	_tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Photo is not a transactional entity")
	}
	ph.config.driver = _tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Photo) String() string {
	var builder strings.Builder
	builder.WriteString("Photo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ph.ID))
	if v := ph.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ph.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ph.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ph.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.Thumbnail; v != nil {
		builder.WriteString("thumbnail=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.TakeTime; v != nil {
		builder.WriteString("take_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ph.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.Team; v != nil {
		builder.WriteString("team=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.Location; v != nil {
		builder.WriteString("location=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ph.Likes; v != nil {
		builder.WriteString("likes=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Photos is a parsable slice of Photo.
type Photos []*Photo

func (ph Photos) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
