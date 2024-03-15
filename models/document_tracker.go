package models

import "time"

type DocumentTracker struct {
	Id                int
	DocumentUUID      string
	ViewedCount       int
	SearchedCount     int
	DocumentCreatedAt time.Time

	DocumentDetail Document
}

var DocumentTrackerTable string = "documents_tracker"
