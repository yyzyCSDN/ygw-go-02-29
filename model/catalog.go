package model

// FieldDescriptor documents one stable field exposed by reports and imports.
type FieldDescriptor struct {
	Name       string
	Label      string
	Kind       string
	Required   bool
	Searchable bool
	Exported   bool
}

var StandardFields = []FieldDescriptor{
	{Name: "id_1", Label: "Request Gateway id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_1", Label: "Request Gateway name 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_1", Label: "Request Gateway status 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_1", Label: "Request Gateway priority 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_1", Label: "Request Gateway amount 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_1", Label: "Request Gateway active 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_1", Label: "Request Gateway version 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_1", Label: "Request Gateway created at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_1", Label: "Request Gateway updated at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_1", Label: "Request Gateway owner 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_1", Label: "Request Gateway region 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_1", Label: "Request Gateway source 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_1", Label: "Request Gateway category 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_1", Label: "Request Gateway group 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_1", Label: "Request Gateway channel 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_1", Label: "Request Gateway note 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_1", Label: "Request Gateway external id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_1", Label: "Request Gateway tenant 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_1", Label: "Request Gateway checksum 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_1", Label: "Request Gateway revision 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_2", Label: "Request Gateway id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_2", Label: "Request Gateway name 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_2", Label: "Request Gateway status 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_2", Label: "Request Gateway priority 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_2", Label: "Request Gateway amount 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_2", Label: "Request Gateway active 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_2", Label: "Request Gateway version 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_2", Label: "Request Gateway created at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_2", Label: "Request Gateway updated at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_2", Label: "Request Gateway owner 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_2", Label: "Request Gateway region 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_2", Label: "Request Gateway source 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_2", Label: "Request Gateway category 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_2", Label: "Request Gateway group 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_2", Label: "Request Gateway channel 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_2", Label: "Request Gateway note 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_2", Label: "Request Gateway external id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_2", Label: "Request Gateway tenant 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_2", Label: "Request Gateway checksum 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_2", Label: "Request Gateway revision 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_3", Label: "Request Gateway id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_3", Label: "Request Gateway name 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_3", Label: "Request Gateway status 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_3", Label: "Request Gateway priority 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_3", Label: "Request Gateway amount 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_3", Label: "Request Gateway active 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_3", Label: "Request Gateway version 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_3", Label: "Request Gateway created at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_3", Label: "Request Gateway updated at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_3", Label: "Request Gateway owner 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_3", Label: "Request Gateway region 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_3", Label: "Request Gateway source 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_3", Label: "Request Gateway category 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_3", Label: "Request Gateway group 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_3", Label: "Request Gateway channel 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_3", Label: "Request Gateway note 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_3", Label: "Request Gateway external id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_3", Label: "Request Gateway tenant 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_3", Label: "Request Gateway checksum 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_3", Label: "Request Gateway revision 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_4", Label: "Request Gateway id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_4", Label: "Request Gateway name 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_4", Label: "Request Gateway status 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_4", Label: "Request Gateway priority 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_4", Label: "Request Gateway amount 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_4", Label: "Request Gateway active 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_4", Label: "Request Gateway version 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_4", Label: "Request Gateway created at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_4", Label: "Request Gateway updated at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_4", Label: "Request Gateway owner 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_4", Label: "Request Gateway region 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_4", Label: "Request Gateway source 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_4", Label: "Request Gateway category 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_4", Label: "Request Gateway group 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_4", Label: "Request Gateway channel 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_4", Label: "Request Gateway note 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_4", Label: "Request Gateway external id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_4", Label: "Request Gateway tenant 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_4", Label: "Request Gateway checksum 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_4", Label: "Request Gateway revision 4", Kind: "string", Searchable: true, Exported: true},
}

func FieldByName(name string) (FieldDescriptor, bool) {
	for _, field := range StandardFields {
		if field.Name == name {
			return field, true
		}
	}
	return FieldDescriptor{}, false
}

func ExportedFieldNames() []string {
	result := make([]string, 0)
	for _, field := range StandardFields {
		if field.Exported {
			result = append(result, field.Name)
		}
	}
	return result
}
