package models

import (
	"time"

	"github.com/pebbe/util"
	check "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

type ExtensionSuite struct {
}

var _ = check.Suite(&ExtensionSuite{})

func (e *ExtensionSuite) TestMarshalStringExtension(c *check.C) {
	ext := &Extension{
		Url:         "http://example.org/fhir/extensions/foo",
		ValueString: "bar",
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "string",
			},
		},
		"foo": "bar",
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalStringExtension(c *check.C) {
	expected := Extension{
		Url:         "http://example.org/fhir/extensions/foo",
		ValueString: "bar",
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "string",
			},
		},
		"foo": "bar",
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestMarshalIntegerExtension(c *check.C) {
	fifty := int32(50)
	ext := &Extension{
		Url:          "http://example.org/fhir/extensions/foo",
		ValueInteger: &fifty,
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "integer",
			},
		},
		"foo": 50,
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalIntegerExtension(c *check.C) {
	fifty := int32(50)
	expected := Extension{
		Url:          "http://example.org/fhir/extensions/foo",
		ValueInteger: &fifty,
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "integer",
			},
		},
		"foo": 50,
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestMarshalBooleanExtension(c *check.C) {
	t := true
	ext := &Extension{
		Url:          "http://example.org/fhir/extensions/foo",
		ValueBoolean: &t,
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "boolean",
			},
		},
		"foo": true,
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalBooleanExtension(c *check.C) {
	t := true
	expected := Extension{
		Url:          "http://example.org/fhir/extensions/foo",
		ValueBoolean: &t,
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "boolean",
			},
		},
		"foo": true,
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestMarshalCodeableConceptExtension(c *check.C) {
	ext := &Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueCodeableConcept: &CodeableConcept{
			Coding: []Coding{
				{System: "http://example.org/fhir/valuesets/foo", Code: "bar"},
				{System: "http://example.org/fhir/valuesets/fooz", Code: "barz"},
			},
			Text: "bar",
		},
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "CodeableConcept",
			},
		},
		"foo": bson.M{
			"coding": []interface{}{
				bson.M{"system": "http://example.org/fhir/valuesets/foo", "code": "bar"},
				bson.M{"system": "http://example.org/fhir/valuesets/fooz", "code": "barz"},
			},
			"text": "bar",
		},
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalCodeableConceptExtension(c *check.C) {
	expected := Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueCodeableConcept: &CodeableConcept{
			Coding: []Coding{
				{System: "http://example.org/fhir/valuesets/foo", Code: "bar"},
				{System: "http://example.org/fhir/valuesets/fooz", Code: "barz"},
			},
			Text: "bar",
		},
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "CodeableConcept",
			},
		},
		"foo": bson.M{
			"coding": []interface{}{
				bson.M{"system": "http://example.org/fhir/valuesets/foo", "code": "bar"},
				bson.M{"system": "http://example.org/fhir/valuesets/fooz", "code": "barz"},
			},
			"text": "bar",
		},
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestMarshalReferenceExtension(c *check.C) {
	t := true
	ext := &Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueReference: &Reference{
			Reference:    "Practitioner/123",
			ReferencedID: "123",
			Type:         "Practitioner",
			External:     &t,
		},
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "Reference",
			},
		},
		"foo": bson.M{
			"reference":   "Practitioner/123",
			"reference__id": "123",
			"reference__type":        "Practitioner",
			"reference__external":    true,
		},
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalReferenceExtension(c *check.C) {
	t := true
	expected := Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueReference: &Reference{
			Reference:    "Practitioner/123",
			ReferencedID: "123",
			Type:         "Practitioner",
			External:     &t,
		},
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "Reference",
			},
		},
		"foo": bson.M{
			"reference":   "Practitioner/123",
			"reference__id": "123",
			"reference__type":        "Practitioner",
			"reference__external":    true,
		},
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestMarshalDateTimeExtension(c *check.C) {
	ext := &Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueDateTime: &FHIRDateTime{
			Time:      time.Date(2012, time.March, 1, 12, 0, 0, 0, time.UTC),
			Precision: Precision(Timestamp),
		},
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "dateTime",
			},
		},
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m["@context"], check.DeepEquals, expected["@context"])
	c.Assert(m["foo"].(bson.M)["__from"].(time.Time).Unix(), check.Equals, time.Date(2012, time.March, 1, 12, 0, 0, 0, time.UTC).Unix())
	c.Assert(m["foo"].(bson.M)["__to"].(time.Time).Unix(), check.Equals, time.Date(2012, time.March, 1, 12, 0, 1, 0, time.UTC).Unix())
	c.Assert(m["foo"].(bson.M)["__strDate"].(string), check.Equals, "2012-03-01T12:00:00Z")
}

func (e *ExtensionSuite) TestUnmarshalDateTimeExtension(c *check.C) {
	expected := Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueDateTime: &FHIRDateTime{
			Time:      time.Date(2012, time.March, 1, 12, 0, 0, 0, time.UTC),
			Precision: Precision(Timestamp),
		},
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "dateTime",
			},
		},
		"foo": time.Date(2012, time.March, 1, 12, 0, 0, 0, time.UTC),
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	// Can't do deep equals of whole object because the time location won't match (despite having the same offset)
	c.Assert(ext.Url, check.Equals, expected.Url)
	c.Assert(ext.ValueDateTime.Precision, check.Equals, expected.ValueDateTime.Precision)
	c.Assert(ext.ValueDateTime.Time.Unix(), check.Equals, expected.ValueDateTime.Time.Unix())
}

func (e *ExtensionSuite) TestMarshalRangeExtension(c *check.C) {
	// l := float64(10)
	// h := float64(20)
	l, err := NewDecimal("10")
	util.CheckErr(err)
	h, err := NewDecimal("20")
	util.CheckErr(err)

	ext := &Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueRange: &Range{
			Low:  &Quantity{Value: l, Unit: "mm"},
			High: &Quantity{Value: h, Unit: "mm"},
		},
	}

	expected := bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "Range",
			},
		},
		"foo": bson.M{
			"low":  bson.M{
				"value": bson.M{
							"__to": float64(10.5),
							"__from": float64(9.5),
							"__num": float64(10),
							"__strNum": "10",
						},
				"unit": "mm"},
			"high":  bson.M{
				"value": bson.M{
							"__to": float64(20.5),
							"__from": float64(19.5),
							"__num": float64(20),
							"__strNum": "20",
						},
				"unit": "mm"},
		},
	}

	// This is where SetBSON is called to marshal it into BSON bytes
	data, err := bson.Marshal(ext)
	util.CheckErr(err)

	// Now unmarshal it back to a map so we can check it against the expected values
	var m bson.M
	err = bson.Unmarshal(data, &m)
	util.CheckErr(err)

	c.Assert(m, check.DeepEquals, expected)
}

func (e *ExtensionSuite) TestUnmarshalRangeExtension(c *check.C) {
	// l := float64(10)
	// h := float64(20)
	l, err := NewDecimal("10")
	util.CheckErr(err)
	h, err := NewDecimal("20")
	util.CheckErr(err)

	expected := Extension{
		Url: "http://example.org/fhir/extensions/foo",
		ValueRange: &Range{
			Low:  &Quantity{Value: l, Unit: "mm"},
			High: &Quantity{Value: h, Unit: "mm"},
		},
	}

	// First marshal the BSON representation into a BSON bytestream
	data, err := bson.Marshal(bson.M{
		"@context": bson.M{
			"foo": bson.M{
				"@id":   "http://example.org/fhir/extensions/foo",
				"@type": "Range",
			},
		},
		"foo": bson.M{
			// "low":  bson.M{"value": float64(10), "unit": "mm"},
			// "high": bson.M{"value": float64(20), "unit": "mm"},
			"low":  bson.M{
				"value": bson.M{
							"__to": float64(10.5),
							"__from": float64(9.5),
							"__num": float64(10),
							"__strNum": "10",
						},
				"unit": "mm"},
			"high":  bson.M{
				"value": bson.M{
							"__to": float64(20.5),
							"__from": float64(19.5),
							"__num": float64(20),
							"__strNum": "20",
						},
				"unit": "mm"},
		},
	})
	util.CheckErr(err)

	// Now unmarshal it into the extension and check it
	var ext Extension
	err = bson.Unmarshal(data, &ext)
	util.CheckErr(err)

	c.Assert(ext, check.DeepEquals, expected)
}
