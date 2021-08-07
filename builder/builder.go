package main

import "errors"

// The NotificationBuilder intended to create final Object
type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

// The constructor with mandatory fields: title, subtitle
func newBuilder(title string, subtitle string) *NotificationBuilder {
	nb := &NotificationBuilder{}
	nb.Title = title
	nb.SubTitle = subtitle
	return nb
}

func (nb *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	nb.Title = title
	return nb
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) *NotificationBuilder {
	nb.SubTitle = subtitle
	return nb
}

func (nb *NotificationBuilder) SetMessage(message string) *NotificationBuilder {
	nb.Message = message
	return nb
}

func (nb *NotificationBuilder) SetImage(image string) *NotificationBuilder {
	nb.Image = image
	return nb
}

func (nb *NotificationBuilder) SetIcon(icon string) *NotificationBuilder {
	nb.Icon = icon
	return nb
}

func (nb *NotificationBuilder) SetPriority(pri int) *NotificationBuilder {
	nb.Priority = pri
	return nb
}

func (nb *NotificationBuilder) SetType(notType string) *NotificationBuilder {
	nb.NotType = notType
	return nb
}

// The Build method returns a fully finished Notification object
func (nb *NotificationBuilder) Build() (*Notification, error) {
	// Error checking can be done at the Build stage
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, errors.New("you have to specify SubTitle when using the icon")
	}

	if nb.Priority < 0 || nb.Priority > 9 {
		return nil, errors.New("priority must be 0..9")
	}

	// Return a newly created Notification object using the current settings
	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
