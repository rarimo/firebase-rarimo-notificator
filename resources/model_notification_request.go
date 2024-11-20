/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NotificationRequest struct {
	// Content of the notification (optional)
	Content *string `json:"content,omitempty"`
	// Description of the notification
	Description string `json:"description"`
	// Target platform for the notification
	Target string `json:"target"`
	// Title of the notification
	Title string `json:"title"`
	// Topic of the notification
	Topic string `json:"topic"`
	// Type of the notification
	Type string `json:"type"`
}
