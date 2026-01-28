package flowconsole


type UserArgs struct {
	Badge *string `field:"optional" json:"badge" yaml:"badge"`
	BelongsTo interface{} `field:"optional" json:"belongsTo" yaml:"belongsTo"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Id *string `field:"optional" json:"id" yaml:"id"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	Tone *string `field:"optional" json:"tone" yaml:"tone"`
	Role *string `field:"optional" json:"role" yaml:"role"`
}

