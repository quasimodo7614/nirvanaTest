package descriptors

import (
	"myproject/pkg/message"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/messages",
		Definitions: []def.Definition{listMessages},
	}, {
		Path:        "/messages/{message}",
		Definitions: []def.Definition{getMessage},
	}, {
		Path:        "/messages",
		Definitions: []def.Definition{createMessage},
	},
		{
			Path:        "/messages",
			Definitions: []def.Definition{updateMessage},
		},
		{
			Path:        "/messages/{message}",
			Definitions: []def.Definition{deleteMessage},
		},
	}...)
}

var listMessages = def.Definition{
	Method:      def.List,
	Summary:     "List Messages",
	Description: "Query a specified number of messages and returns an array",
	Function:    message.ListMessages,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "count",
			Default:     10,
			Description: "Number of messages",
		},
	},
	Results: def.DataErrorResults("A list of messages"),
}

var getMessage = def.Definition{
	Method:      def.Get,
	Summary:     "Get Message",
	Description: "Get a message by id",
	Function:    message.GetMessage,
	Parameters: []def.Parameter{
		def.PathParameterFor("message", "Message id"),
	},
	Results: def.DataErrorResults("A message"),
}

var createMessage = def.Definition{
	Method:      def.Create,
	Summary:     "Create Message",
	Description: "Create a message ",
	Function:    message.CreateMessage,
	Parameters: []def.Parameter{
		def.BodyParameterFor("Create Message request"),
	},
	Results: def.DataErrorResults("Result of Create"),
}

var updateMessage = def.Definition{
	Method:      def.Update,
	Summary:     "Update Message",
	Description: "Update a message ",
	Function:    message.UpdateMessage,
	Parameters: []def.Parameter{
		def.BodyParameterFor("Update Message request"),
	},
	Results: def.DataErrorResults("Result of Update"),
}

var deleteMessage = def.Definition{
	Method:      def.Delete,
	Summary:     "Delete Message",
	Description: "Delete a message ",
	Function:    message.DeleteMessage,
	Parameters: []def.Parameter{
		def.PathParameterFor("message", "message id"),
	},
	Results: []def.Result{def.ErrorResult()},
}
