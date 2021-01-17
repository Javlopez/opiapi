package domain

type Point struct {
	CartoDBID int    `gorm:"primary_key" json:"id"`
	TheGeom   string `gorm:"type:varchar(50)" json:"the_geom"`
	Type      string `gorm:"type:varchar(50)" json:"type"`
	Latitude  string `gorm:"type:varchar(20)" json:"latitude"`
	Longitude string `gorm:"type:varchar(20)" json:"longitude"`
	Color     string `gorm:"type:varchar(10)" json:"color"`
}

//PointRepository interface
type PointRepository interface {
	SavePoints(points []Point) error
	Fetch() ([]Point, error)
}
