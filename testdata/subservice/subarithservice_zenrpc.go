// Code generated by zenrpc; DO NOT EDIT.

package subarithservice

import (
	"context"
	"encoding/json"

	"github.com/semrush/zohar-soul/v2"
	"github.com/semrush/zohar-soul/v2/smd"

	"github.com/semrush/zohar-soul/v2/testdata/model"
)

var RPC = struct {
	SubArithService struct{ Sum, Positive, ReturnPointFromSamePackage, GetPoints, GetPointsFromSamePackage, DoSomethingWithPoint, Multiply, CheckError, CheckZenRPCError, Divide, Pow, Pi, SumArray string }
}{
	SubArithService: struct{ Sum, Positive, ReturnPointFromSamePackage, GetPoints, GetPointsFromSamePackage, DoSomethingWithPoint, Multiply, CheckError, CheckZenRPCError, Divide, Pow, Pi, SumArray string }{
		Sum:                        "sum",
		Positive:                   "positive",
		ReturnPointFromSamePackage: "returnpointfromsamepackage",
		GetPoints:                  "getpoints",
		GetPointsFromSamePackage:   "getpointsfromsamepackage",
		DoSomethingWithPoint:       "dosomethingwithpoint",
		Multiply:                   "multiply",
		CheckError:                 "checkerror",
		CheckZenRPCError:           "checkzenrpcerror",
		Divide:                     "divide",
		Pow:                        "pow",
		Pi:                         "pi",
		SumArray:                   "sumarray",
	},
}

func (SubArithService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Sum": {
				Description: `Sum sums two digits and returns error with error code as result and IP from context.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Boolean,
				},
			},
			"Positive": {
				Description: ``,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Boolean,
				},
			},
			"ReturnPointFromSamePackage": {
				Description: ``,
				Parameters: []smd.JSONSchema{
					{
						Name:        "p",
						Optional:    false,
						Description: ``,
						Type:        smd.Object,
						Properties: map[string]smd.Property{
							"Name": {
								Description: ``,
								Type:        smd.String,
							},
							"SomeField": {
								Description: ``,
								Type:        smd.String,
							},
							"Measure": {
								Description: ``,
								Type:        smd.Float,
							},
							"A": {
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							"B": {
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							"when": {
								Description: `when it happened`,
								Ref:         "#/definitions/time.Time",
								Type:        smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"time.Time": {
								Type:       "object",
								Properties: map[string]smd.Property{},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"Name": {
							Description: ``,
							Type:        smd.String,
						},
						"SomeField": {
							Description: ``,
							Type:        smd.String,
						},
						"Measure": {
							Description: ``,
							Type:        smd.Float,
						},
						"A": {
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						"B": {
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						"when": {
							Description: `when it happened`,
							Ref:         "#/definitions/time.Time",
							Type:        smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"time.Time": {
							Type:       "object",
							Properties: map[string]smd.Property{},
						},
					},
				},
			},
			"GetPoints": {
				Description: ``,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"$ref": "#/definitions/model.Point",
					},
					Definitions: map[string]smd.Definition{
						"model.Point": {
							Type: "object",
							Properties: map[string]smd.Property{
								"Name": {
									Description: ``,
									Type:        smd.String,
								},
								"SomeField": {
									Description: ``,
									Type:        smd.String,
								},
								"Measure": {
									Description: ``,
									Type:        smd.Float,
								},
								"X": {
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								"Y": {
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								"ConnectedObject": {
									Description: ``,
									Ref:         "#/definitions/objects.AbstractObject",
									Type:        smd.Object,
								},
							},
						},
						"objects.AbstractObject": {
							Type: "object",
							Properties: map[string]smd.Property{
								"Name": {
									Description: ``,
									Type:        smd.String,
								},
								"SomeField": {
									Description: ``,
									Type:        smd.String,
								},
								"Measure": {
									Description: ``,
									Type:        smd.Float,
								},
							},
						},
					},
				},
			},
			"GetPointsFromSamePackage": {
				Description: ``,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"$ref": "#/definitions/Point",
					},
					Definitions: map[string]smd.Definition{
						"Point": {
							Type: "object",
							Properties: map[string]smd.Property{
								"Name": {
									Description: ``,
									Type:        smd.String,
								},
								"SomeField": {
									Description: ``,
									Type:        smd.String,
								},
								"Measure": {
									Description: ``,
									Type:        smd.Float,
								},
								"A": {
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								"B": {
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								"when": {
									Description: `when it happened`,
									Ref:         "#/definitions/time.Time",
									Type:        smd.Object,
								},
							},
						},
						"time.Time": {
							Type:       "object",
							Properties: map[string]smd.Property{},
						},
					},
				},
			},
			"DoSomethingWithPoint": {
				Description: ``,
				Parameters: []smd.JSONSchema{
					{
						Name:        "p",
						Optional:    false,
						Description: ``,
						Type:        smd.Object,
						Properties: map[string]smd.Property{
							"Name": {
								Description: ``,
								Type:        smd.String,
							},
							"SomeField": {
								Description: ``,
								Type:        smd.String,
							},
							"Measure": {
								Description: ``,
								Type:        smd.Float,
							},
							"X": {
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							"Y": {
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							"ConnectedObject": {
								Description: ``,
								Ref:         "#/definitions/objects.AbstractObject",
								Type:        smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"objects.AbstractObject": {
								Type: "object",
								Properties: map[string]smd.Property{
									"Name": {
										Description: ``,
										Type:        smd.String,
									},
									"SomeField": {
										Description: ``,
										Type:        smd.String,
									},
									"Measure": {
										Description: ``,
										Type:        smd.Float,
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"Name": {
							Description: ``,
							Type:        smd.String,
						},
						"SomeField": {
							Description: ``,
							Type:        smd.String,
						},
						"Measure": {
							Description: ``,
							Type:        smd.Float,
						},
						"X": {
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						"Y": {
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						"ConnectedObject": {
							Description: ``,
							Ref:         "#/definitions/objects.AbstractObject",
							Type:        smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"objects.AbstractObject": {
							Type: "object",
							Properties: map[string]smd.Property{
								"Name": {
									Description: ``,
									Type:        smd.String,
								},
								"SomeField": {
									Description: ``,
									Type:        smd.String,
								},
								"Measure": {
									Description: ``,
									Type:        smd.Float,
								},
							},
						},
					},
				},
			},
			"Multiply": {
				Description: `Multiply multiples two digits and returns result.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Integer,
				},
			},
			"CheckError": {
				Description: `CheckError throws error is isErr true.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "isErr",
						Optional:    false,
						Description: ``,
						Type:        smd.Boolean,
					},
				},
				Errors: map[int]string{
					500: "test error",
				},
			},
			"CheckZenRPCError": {
				Description: `CheckError throws zenrpc error is isErr true.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "isErr",
						Optional:    false,
						Description: ``,
						Type:        smd.Boolean,
					},
				},
				Errors: map[int]string{
					500: "test error",
				},
			},
			"Divide": {
				Description: `Divide divides two numbers.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: `the a`,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: `the b`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"Quo": {
							Description: `Quo docs`,
							Type:        smd.Integer,
						},
						"rem": {
							Description: `Rem docs`,
							Type:        smd.Integer,
						},
					},
				},
				Errors: map[int]string{
					401:    "we do not serve 1",
					-32603: "divide by zero",
				},
			},
			"Pow": {
				Description: `Pow returns x**y, the base-x exponential of y. If Exp is not set then default value is 2.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "base",
						Optional:    false,
						Description: ``,
						Type:        smd.Float,
					},
					{
						Name:        "exp",
						Optional:    true,
						Description: `exponent could be empty`,
						Type:        smd.Float,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Float,
				},
			},
			"Pi": {
				Description: `PI returns math.Pi.`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Float,
				},
			},
			"SumArray": {
				Description: `SumArray returns sum all items from array`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "array",
						Optional:    true,
						Description: ``,
						Type:        smd.Array,
						Items: map[string]string{
							"type": smd.Float,
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Float,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s SubArithService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.SubArithService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Sum(ctx, args.A, args.B))

	case RPC.SubArithService.Positive:
		resp.Set(s.Positive())

	case RPC.SubArithService.ReturnPointFromSamePackage:
		var args = struct {
			P Point `json:"p"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"p"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.ReturnPointFromSamePackage(args.P))

	case RPC.SubArithService.GetPoints:
		resp.Set(s.GetPoints())

	case RPC.SubArithService.GetPointsFromSamePackage:
		resp.Set(s.GetPointsFromSamePackage())

	case RPC.SubArithService.DoSomethingWithPoint:
		var args = struct {
			P model.Point `json:"p"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"p"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.DoSomethingWithPoint(args.P))

	case RPC.SubArithService.Multiply:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Multiply(args.A, args.B))

	case RPC.SubArithService.CheckError:
		var args = struct {
			IsErr bool `json:"isErr"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"isErr"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.CheckError(args.IsErr))

	case RPC.SubArithService.CheckZenRPCError:
		var args = struct {
			IsErr bool `json:"isErr"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"isErr"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.CheckZenRPCError(args.IsErr))

	case RPC.SubArithService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Divide(args.A, args.B))

	case RPC.SubArithService.Pow:
		var args = struct {
			Base float64  `json:"base"`
			Exp  *float64 `json:"exp"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"base", "exp"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:exp=2 	exponent could be empty
		if args.Exp == nil {
			var v float64 = 2
			args.Exp = &v
		}

		resp.Set(s.Pow(args.Base, args.Exp))

	case RPC.SubArithService.Pi:
		resp.Set(s.Pi())

	case RPC.SubArithService.SumArray:
		var args = struct {
			Array *[]float64 `json:"array"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"array"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:array=[]float64{1,2,4}
		if args.Array == nil {
			var v []float64 = []float64{1, 2, 4}
			args.Array = &v
		}

		resp.Set(s.SumArray(args.Array))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
