package ingestion

const (
	index   = "omniscience"
	mapping = `
		{
			"settings":{
				"number_of_shards": 1,
				"number_of_replicas": 0
			},
			"mappings":{
				"document":{
					"properties":{
						"id":{
							"type":"keyword"
						},
						"name":{
							"type":"text"
						},
						"description":{
							"type":"text"
						},
						"service":{
							"type":"text"
						},
						"content":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
						"url":{
							"type":"text"
						},
						"last_modified":{
							"type":"date"
						},
						"suggest_field":{
							"type":"completion"
						}
					}
				}
			}
		}`
)
