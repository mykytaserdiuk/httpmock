package generator_test

//mock := &models.MockScheme{
//	Port: ":8080",
//	Paths: []*models.Path{
//		{
//			Path: "/login/{user_id}",
//			Endpoints: models.Endpoints{
//				{
//					Parameters: []*models.Parameter{{
//						In:          "query",
//						Placeholder: "user_id",
//						Value:       "11",
//					}},
//					Method: models.MethodGet,
//					Response: models.Response{
//						Header: map[string][]string{"Agent": {"linux"}},
//						Type:   "application/json",
//						Status: http.StatusOK,
//						Body:   "O KEY",
//					},
//					Request: models.Request{
//						Header:   map[string][]string{"Agent": {"Chrome"}},
//						Expected: "55",
//						Type:     "application/json",
//					},
//				},
//			}},
//	},
//}
