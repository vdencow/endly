package common

import (
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/ssh"
)





type Map map[string]interface{}
func (s *Map) Put(key string, value interface{}) {
	(*s)[key] = value
}


func (s *Map) Has(key string) bool {
	 _, found := (*s)[key];
	return found
}


func (s *Map) Get(key string) interface{} {
	if result, found := (*s)[key]; found {
		return result
	}
	return nil

}

func (s *Map) GetString(key string) string {
	if result, found := (*s)[key]; found {
		return toolbox.AsString(result)
	}
	return ""
}


func (s *Map) GetCommandSession(key string) *ssh.MultiCommandSession {
	if result, found := (*s)[key]; found {
		if session, ok := result.(*ssh.MultiCommandSession); ok {
			return session
		}
	}
	return nil
}

func (s *Map) GetInt(key string) int {
	if result, found := (*s)[key]; found {
		return toolbox.AsInt(result)
	}
	return 0
}

func (s *Map) GetFloat(key string) float64 {
	if result, found := (*s)[key]; found {
		return toolbox.AsFloat(result)
	}
	return 0.0
}

func (s *Map) GetBoolean(key string) bool {
	if result, found := (*s)[key]; found {
		return toolbox.AsBoolean(result)
	}
	return false
}


func (s *Map) GetCollection(key string) *Collection {
	if result, found := (*s)[key]; found {
		collection, ok := result.(*Collection)
		if ok {
			return collection
		}
	}
	return nil

}


func (s *Map) GetMap(key string) Map {
	if result, found := (*s)[key]; found {
		aMap, ok := result.(Map)
		if ok {
			return aMap
		}
	}
	return nil

}



func NewMap() Map {
	return make(map[string]interface{})
}
