## **基础** **URL**

```Plain
https://api.recruitment.kkkstra.cn/api/v1/
```

## **用户管理**

### **1. 注册新用户**

**HTTP 方法：** `POST`  

**资源路径：** `/user`  

**描述：** 注册新用户（公司或求职者）。

**请求体：**

```JSON
{
    "username": "test",
    "email": "test@test.com",
    "password": "123456",
    "role": 1,
    "age": 22,
    "degree": 1
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 3
    },
    "msg": "register success"
}
```

- **201 Created**: 用户注册成功。
- **400 Bad Request**: 请求参数无效或缺失。

### **2. 用户登录**

**HTTP 方法：** `POST`  

**资源路径：** `/session`  

**描述：** 用户登录并获取 JWT 令牌。

**请求体：**

```JSON
{
    "username": "test",
    "password": "123456"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "expire": 168,
        "id": 1,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoxLCJleHAiOjE3MzAxODM4NzUsImlhdCI6MTczMDE4MzcwNywiaXNzIjoiSFVTVCIsInN1YiI6IjEifQ.s3NYoVcAIkvgUAQRBHJ47cfSVJNu5N_qYo9aEIEY7No"
    },
    "msg": "login success"
}
```

- **200 OK**: 登录成功并返回 JWT。
- **401 Unauthorized**: 登录失败，邮箱或密码错误。

### **3. 获取用户信息**

**HTTP 方法：** `GET`  

**资源路径：** `/user/:id/profile`

**描述：** 获取当前登录用户的详细信息。

**请求头：**

```JSON
{
    "Authorization": "Bearer <your_jwt_token>"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "profile": {
            "id": 1,
            "username": "test",
            "email": "test@test.com",
            "role": 1,
            "age": 22,
            "degree": 1
        }
    },
    "msg": "get user profile success"
}
```

- **200 OK**: 返回当前用户信息。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。

### **4. 更新当前用户信息**

**HTTP 方法：** `PUT`  

**资源路径：** `/user/:id/profile`  

**描述：** 更新当前用户的信息（如用户名或邮箱）。

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "username": "test",
    "email": "test@test.com",
    "age": 22,
    "degree": 1
}
```

**响应：**

- **200 OK**: 用户信息更新成功。
- **400 Bad Request**: 请求参数无效。
- **401 Unauthorized**: 无效的 JWT 令牌。

## **职位管理**

### **1. 创建职位**

**HTTP 方法：** `POST`  

**资源路径：** `/jobs`  

**描述：** 创建新职位（需要公司用户身份验证）。

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "title": "算法工程师",
    "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
    "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
    "location": "上海",
    "company": "比特跳动",
    "salary": "40k*15",
    "job_type": "研发"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 1
    },
    "msg": "create job success"
}
```

- **201 Created**: 职位创建成功。
- **400 Bad Request**: 请求参数无效或缺失。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。

### **2. 获取所有职位**

**HTTP 方法：** `GET`  

**资源路径：** `/jobs/?own={True/False}`  

**描述：** 获取所有职位的列表。

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "jobs": [
            {
                "id": 1,
                "title": "后端开发工程师",
                "description": "TikTok研发团队，旨在实现TikTok业务的研发工作，搭建及维护业界领先的产品。加入我们，你能接触到包括用户增长、社交、直播、电商C端、内容创造、内容消费等核心业务场景，支持产品在全球赛道上高速发展；也能接触到包括服务架构、基础技术等方向上的技术挑战，保障业务持续高质量、高效率、且安全地为用户服务；同时还能为不同业务场景提供全面的技术解决方案，优化各项产品指标及用户体验。\n在这里， 有大牛带队与大家一同不断探索前沿， 突破想象空间。 在这里，你的每一行代码都将服务亿万用户。在这里，团队专业且纯粹，合作氛围平等且轻松。目前在北京，上海，杭州、广州、深圳分别开放多个岗位机会。",
                "demand": "1、2025届获得本科及以上学历，计算机、软件、电子信息等相关专业；\n2、热爱计算机科学和互联网技术，精通至少一门编程语言，包括但不仅限于：Java、C、C++、PHP、 Python、Go；\n3、掌握扎实的计算机基础知识，深入理解数据结构、算法和操作系统知识；\n4、有优秀的逻辑分析能力，能够对业务逻辑进行合理的抽象和拆分；\n5、有强烈的求知欲，优秀的学习和沟通能力。\n",
                "location": "上海",
                "company": "比特跳动",
                "salary": "24k*15",
                "job_type": "研发",
                "owner_id": 1
            },
            {
                "id": 2,
                "title": "算法工程师",
                "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
                "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
                "location": "上海",
                "company": "比特跳动",
                "salary": "40k*15",
                "job_type": "研发-算法",
                "owner_id": 1
            }
        ]
    },
    "msg": "get all jobs success"
}
```

- **200 OK**: 返回职位列表。

### **3. 获取单个职位详情**

**HTTP 方法：** `GET`  

**资源路径：** `/jobs/:id`  

**描述：** 获取指定职位的详细信息。

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "job": {
            "id": 1,
            "title": "算法工程师",
            "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
            "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
            "location": "上海",
            "company": "比特跳动",
            "salary": "40k*15",
            "job_type": "研发",
            "owner_id": 1
        }
    },
    "msg": "get job success"
}
```

- **200 OK**: 返回职位详情。
- **404 Not Found**: 找不到该职位。

### **4. 更新职位信息**

**HTTP 方法：** `PUT`  

**资源路径：** `/jobs/:id`  

**描述：** 更新指定职位的信息（需要公司用户身份验证）。

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "title": "算法工程师",
    "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
    "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
    "location": "上海",
    "company": "比特跳动",
    "salary": "40k*15",
    "job_type": "研发"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 7
    },
    "msg": "create job success"
}
```

- **200 OK**: 职位信息更新成功。
- **400 Bad Request**: 请求参数无效或缺失。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。
- **404 Not Found**: 找不到该职位。

### **5. 删除职位**

**HTTP 方法：** `DELETE`  

**资源路径：** `/jobs/:id`  

**描述：** 删除指定职位（需要公司用户身份验证）。

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应：**

- **204 No Content**: 职位删除成功，无内容返回。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。
- **404 Not Found**: 找不到该职位。

## **简历管理**

### **1. 新建简历**

**HTTP 方法：** `POST`  

**资源路径：** `/resumes`  

**描述：** 上传求职者的简历文件（支持 PDF 或 DOC 格式）。

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```Plain
{
  "name": "John Doe",
  "gender": 1,
  "phone": "123-456-7890",
  "email": "john.doe@example.com",
  "wechat": "johnwechat",
  "state": 4,
  "description": "A software engineer with a passion for developing innovative solutions.",
  "education": [
    {
      "school": "University of Example",
      "major": "Computer Science",
      "degree": 1,
      "start_time": "2015-09-01T00:00:00Z",
      "end_time": "2019-06-01T00:00:00Z"
    }
  ],
  "experience": [
    {
      "company": "Tech Solutions",
      "position": "Software Developer",
      "start_time": "2020-01-01T00:00:00Z",
      "end_time": "2022-12-31T00:00:00Z"
    }
  ],
  "project": [
    {
      "name": "Project Alpha",
      "description": "Developed a web application to optimize workflows.",
      "start_time": "2021-03-01T00:00:00Z",
      "end_time": "2021-08-01T00:00:00Z"
    }
  ]
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 3
    },
    "msg": "create resume success"
}
```

- **200 OK**: 简历上传成功。
- **400 Bad Request**: 文件格式无效或缺失文件。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。

1. ### 根据user_id获取简历

**HTTP 方法：** `GET`  

**资源路径：** `/resumes/:user_id`  

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "resume": {
            "id": 1,
            "user_id": 3,
            "name": "John Doe",
            "gender": 1,
            "phone": "123-456-7890",
            "email": "john.doe@example.com",
            "wechat": "johnwechat",
            "state": 4,
            "description": "A software engineer with a passion for developing innovative solutions.",
            "education": [
                {
                    "school": "University of Example",
                    "major": "Computer Science",
                    "degree": 1,
                    "start_time": "2015-09-01T00:00:00Z",
                    "end_time": "2019-06-01T00:00:00Z"
                }
            ],
            "experience": [
                {
                    "company": "Tech Solutions",
                    "position": "Software Developer",
                    "start_time": "2020-01-01T00:00:00Z",
                    "end_time": "2022-12-31T00:00:00Z"
                }
            ],
            "project": [
                {
                    "name": "Project Alpha",
                    "description": "Developed a web application to optimize workflows.",
                    "start_time": "2021-03-01T00:00:00Z",
                    "end_time": "2021-08-01T00:00:00Z"
                }
            ]
        }
    },
    "msg": "get resume success"
}
```

- **200 OK**: 简历上传成功。
- **400 Bad Request**: 文件格式无效或缺失文件。
- **401 Unauthorized**: 未授权，无效的 JWT 令牌。

1. ### 更新简历信息

**HTTP 方法：** `PUT`  

**资源路径：** `/resumes`  

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
  "name": "John Doe",
  "gender": 1,
  "phone": "123-456-789",
  "email": "john.doe@example.com",
  "wechat": "johnwechat",
  "state": 4,
  "description": "A software engineer with a passion for developing innovative solutions.",
  "education": [
    {
      "school": "College of Example",
      "major": "Computer Science",
      "degree": 1,
      "start_time": "2015-09-01T00:00:00Z",
      "end_time": "2019-06-01T00:00:00Z"
    }
  ],
  "experience": [
    {
      "company": "Tech Solutions",
      "position": "Software Developer",
      "start_time": "2020-01-01T00:00:00Z",
      "end_time": "2022-12-31T00:00:00Z"
    }
  ],
  "project": [
    {
      "name": "Project Alpha",
      "description": "Developed a web application to optimize workflows.",
      "start_time": "2021-03-01T00:00:00Z",
      "end_time": "2021-08-01T00:00:00Z"
    }
  ]
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 1
    },
    "msg": "update resume success"
}
```

## 岗位申请管理

1. ### 用户申请岗位 

**HTTP 方法：** `POST`  

**资源路径：** `/applications`  

**要求用户为 candidate**

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "job_id": 1
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 1
    },
    "msg": "application created"
}
```

1. ### 获取用户的申请列表

**HTTP 方法：** `GET`  

**资源路径：**`/applications`  

**要求用户为 candidate**

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "applications": [
            {
                "id": 2,
                "user_id": 3,
                "job_id": 2,
                "progress": 1
            }
        ]
    },
    "msg": "applications retrieved"
}
```

1. ### 获取岗位的申请者

**HTTP 方法：** `GET`  

**资源路径：** `/applications/job/:job_id`

**要求用户为  job 的 owner**

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "applications": [
            {
                "id": 2,
                "user_id": 3,
                "job_id": 2,
                "progress": 1
            }
        ]
    },
    "msg": "applications retrieved"
}
```

1. ### 更新申请状态 

**HTTP 方法：** `PUT`  

**资源路径：** `/applications/:id`  

**说明：**

```JSON
CandidateApplied  = 1    // candidate 提交申请
RecruiterReviewed = 2    // recruiter 接受申请
// 面试，沟通，etc……
RecruiterAccepted = 3    // recruiter 发出offer
RecruiterRejected = 4    // recruiter 拒绝发出 offer
CandidateAccepted = 5    // candidate 接受offer
CandidateRejected = 6    // candidate 拒绝offer
```

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "id": 2,
    "accepted": false
}
```

**响应：**

```JSON
{
    "code": 0,
    "data": {
        "id": 2,
        "job_id": 2,
        "progress": 4,
        "user_id": 3
    },
    "msg": "progress updated"
}
```

## AI 服务

### **1. 根据简历推荐合适的岗位**

**接口描述**

根据候选人id推荐岗位

- **URL**: `/recommend/jobs`
- **请求方式**: `GET`

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应**

根据推荐从高到低返回岗位id

```JSON
{
    "job": [6, 1, 2]
}
```

### **2.  根据描述推荐合适的岗位**

- **URL**: `/recommend/jobs`
- **请求方式**: `POST`

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**请求体：**

```JSON
{
    "description": "钱多事少离家近"
}
```

**响应：**

根据推荐从高到低返回岗位id

```JSON
{
    "job": [6, 1, 2]
}
```

### **3.  根据岗位描述对候选人简历进行打分和筛选**

- **URL**: `/recommend/resumes?job_id={job_id}`
- **请求方式**: `GET`

**请求头：**

```JSON
{
  "Authorization": "Bearer <your_jwt_token>"
}
```

**响应**

返回每个候选人的得分

```JSON
{
    "score": [
        {"id": 1, "score": 80},
        {"id": 3, "score": 70}
    ]
}
```

## **错误处理**

### 错误码