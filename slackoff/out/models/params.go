package models

type Params struct {
	Template       	string `json:"template"`
	TemplateFile   	string `json:"template_file"`
	FileVars  		 	map[string]string `json:"file_vars"`
	Vars       		 	map[string]string `json:"vars"`
	Channel				 	string `json:"channel"`
	ChannelAppend  	string `json:"channel_append"`
	ChannelFile			string `json:"channel_file"`
	IconUrl					string `json:"icon_url"`
	IconEmoji				string `json:"icon_emoji"`
}
