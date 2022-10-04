package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/config"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/vo"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/client"
	"time"

	pb "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/transport/grpc"
)

type EnumDemoImpl struct {
	conf       *config.Config
	grpcClient pb.HelloworldServiceClient
	restClient *client.HelloworldClient
}

func (receiver *EnumDemoImpl) GetKeyboard(ctx context.Context, layout vo.KeyboardLayout) (data string, err error) {
	//return layout.StringGetter(), nil
	jsonStr := `{
    "data": {
        "colTableHeaders": [
            {
                "data": "",
                "children": [
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "总计",
                                "children": [
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "_______居家养老服务商",
                        "children": [
                            {
                                "data": "1",
                                "children": [
                                    {
                                        "data": "A",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "2",
                                "children": [
                                    {
                                        "data": "B",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "3",
                                "children": [
                                    {
                                        "data": "C",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "4",
                                "children": [
                                    {
                                        "data": "D",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "5",
                                "children": [
                                    {
                                        "data": "E",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "_______街道养老照料中心",
                        "children": [
                            {
                                "data": "1",
                                "children": [
                                    {
                                        "data": "F",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "2",
                                "children": [
                                    {
                                        "data": "G",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "3",
                                "children": [
                                    {
                                        "data": "H",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "4",
                                "children": [
                                    {
                                        "data": "I",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "5",
                                "children": [
                                    {
                                        "data": "J",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "_______社区/村养老服务驿站",
                        "children": [
                            {
                                "data": "1",
                                "children": [
                                    {
                                        "data": "K",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "2",
                                "children": [
                                    {
                                        "data": "L",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "3",
                                "children": [
                                    {
                                        "data": "M",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "4",
                                "children": [
                                    {
                                        "data": "N",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "5",
                                "children": [
                                    {
                                        "data": "O",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "_______养老院",
                        "children": [
                            {
                                "data": "1",
                                "children": [
                                    {
                                        "data": "P",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "2",
                                "children": [
                                    {
                                        "data": "Q",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "3",
                                "children": [
                                    {
                                        "data": "R",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "4",
                                "children": [
                                    {
                                        "data": "S",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "5",
                                "children": [
                                    {
                                        "data": "T",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "_______其他（请说明）__________________",
                        "children": [
                            {
                                "data": "1",
                                "children": [
                                    {
                                        "data": "U",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "2",
                                "children": [
                                    {
                                        "data": "V",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "3",
                                "children": [
                                    {
                                        "data": "W",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "4",
                                "children": [
                                    {
                                        "data": "X",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "5",
                                "children": [
                                    {
                                        "data": "Y",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    }
                ],
                "level": 0
            }
        ],
        "rowTableHeaders": [
            {
                "data": "如果您有养老服务需求，请写出选择以下养老服务（机构）的优先顺序：",
                "children": [
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "Base",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "_______居家养老服务商",
                                "children": [
                                    {
                                        "data": "1",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "2",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "3",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "4",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "5",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "总计",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "平均值",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "Base",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "_______街道养老照料中心",
                                "children": [
                                    {
                                        "data": "1",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "2",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "3",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "4",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "5",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "总计",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "平均值",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "Base",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "_______社区/村养老服务驿站",
                                "children": [
                                    {
                                        "data": "1",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "2",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "3",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "4",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "5",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "总计",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "平均值",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "Base",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "_______养老院",
                                "children": [
                                    {
                                        "data": "1",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "2",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "3",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "4",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "5",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "总计",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "平均值",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    },
                    {
                        "data": "",
                        "children": [
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "Base",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "_______其他（请说明）__________________",
                                "children": [
                                    {
                                        "data": "1",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "2",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "3",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "4",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "5",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "总计",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            },
                            {
                                "data": "",
                                "children": [
                                    {
                                        "data": "平均值",
                                        "children": null,
                                        "level": 3
                                    },
                                    {
                                        "data": "",
                                        "children": null,
                                        "level": 3
                                    }
                                ],
                                "level": 2
                            }
                        ],
                        "level": 1
                    }
                ],
                "level": 0
            }
        ],
        "cells": null
    }
}`
	return jsonStr, nil
}
func (receiver *EnumDemoImpl) GetKeyboard2(ctx context.Context, layout *vo.KeyboardLayout) (data string, err error) {
	return layout.StringGetter(), nil
}
func (receiver *EnumDemoImpl) GetKeyboards(ctx context.Context, layout []vo.KeyboardLayout) (data []string, err error) {
	for _, item := range layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}
func (receiver *EnumDemoImpl) GetKeyboards2(ctx context.Context, layout *[]vo.KeyboardLayout) (data []string, err error) {
	for _, item := range *layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}
func (receiver *EnumDemoImpl) GetKeyboards5(ctx context.Context, layout ...vo.KeyboardLayout) (data []string, err error) {
	for _, item := range layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}

func NewEnumDemo(conf *config.Config, grpcClient pb.HelloworldServiceClient, restClient *client.HelloworldClient) EnumDemo {
	return &EnumDemoImpl{
		conf:       conf,
		grpcClient: grpcClient,
		restClient: restClient,
	}
}

func (receiver *EnumDemoImpl) Keyboard(ctx context.Context, keyboard vo.Keyboard) (data string, err error) {
	return keyboard.Layout.StringGetter(), nil
}

func (receiver *EnumDemoImpl) Greeting(ctx context.Context, greeting string) (data string, err error) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := receiver.grpcClient.GreetingRpc(ctx, &pb.GreetingRpcRequest{Greeting: greeting})
	if err != nil {
		panic(err)
	}
	return r.Data, nil
}

func (receiver *EnumDemoImpl) Greeting1(ctx context.Context, greeting string) (data string, err error) {
	_, data, err = receiver.restClient.Greeting(ctx, nil, greeting)
	return
}
