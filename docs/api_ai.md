## **服务概述**

本服务提供两大功能：

1. **求职者根据简历推荐合适的岗位：通过上传 简历，系统分析简历内容并推荐合适的岗位。**
2. **求职者根据描述推荐合适的岗位：通过描述，系统推荐合适的岗位。**
3. **招聘者根据岗位描述筛选简历并打分**：招聘者提供岗位描述并批量上传候选人的 PDF 简历，系统根据岗位需求打分候选人并排序。

## **1. 根据简历推荐合适的岗位**

### **接口描述**

通过上传求职者的简历 PDF 文件，系统分析简历内容，推荐与简历相匹配的工作岗位。

- **URL**: `/recommend_jobs/resume`
- **请求方式**: `POST`

### 请求体

```JSON
{
    "resume": {
            "id": 1,
            "user_id": 3,
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
    },
    "jobs": [
            {
                "id": 1,
                "title": "算法工程师",
                "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
                "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
                "location": "上海",
                "company": "比特跳动",
                "salary": "40k*15",
                "job_type": "研发",
                "owner_id": 1
            },
            {
                "id": 2,
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
                "id": 6,
                "title": "test",
                "description": "fun",
                "demand": "interest",
                "location": "earth",
                "company": "dancing",
                "salary": "40k*15",
                "job_type": "haha",
                "owner_id": 1
            }
    ]
}
```

### 响应

根据推荐从高到低返回岗位id

```JSON
{
    "job": [6, 1, 2]
}
```

## **2.  根据描述推荐合适的岗位**

### **接口描述**

通过上传求职者的简历 PDF 文件，系统分析简历内容，推荐与简历相匹配的工作岗位。

- **URL**: `/recommend_job/description`
- **请求方式**: `POST`

### 请求体

```JSON
{
    "description": "钱多事少离家近",
    "jobs": [
            {
                "id": 1,
                "title": "算法工程师",
                "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
                "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
                "location": "上海",
                "company": "比特跳动",
                "salary": "40k*15",
                "job_type": "研发"
            },
            {
                "id": 2,
                "title": "后端开发工程师",
                "description": "TikTok研发团队，旨在实现TikTok业务的研发工作，搭建及维护业界领先的产品。加入我们，你能接触到包括用户增长、社交、直播、电商C端、内容创造、内容消费等核心业务场景，支持产品在全球赛道上高速发展；也能接触到包括服务架构、基础技术等方向上的技术挑战，保障业务持续高质量、高效率、且安全地为用户服务；同时还能为不同业务场景提供全面的技术解决方案，优化各项产品指标及用户体验。\n在这里， 有大牛带队与大家一同不断探索前沿， 突破想象空间。 在这里，你的每一行代码都将服务亿万用户。在这里，团队专业且纯粹，合作氛围平等且轻松。目前在北京，上海，杭州、广州、深圳分别开放多个岗位机会。",
                "demand": "1、2025届获得本科及以上学历，计算机、软件、电子信息等相关专业；\n2、热爱计算机科学和互联网技术，精通至少一门编程语言，包括但不仅限于：Java、C、C++、PHP、 Python、Go；\n3、掌握扎实的计算机基础知识，深入理解数据结构、算法和操作系统知识；\n4、有优秀的逻辑分析能力，能够对业务逻辑进行合理的抽象和拆分；\n5、有强烈的求知欲，优秀的学习和沟通能力。\n",
                "location": "上海",
                "company": "比特跳动",
                "salary": "24k*15",
                "job_type": "研发"
            },
            {
                "id": 6,
                "title": "test",
                "description": "fun",
                "demand": "interest",
                "location": "earth",
                "company": "dancing",
                "salary": "40k*15",
                "job_type": "haha"
            }
    ]
}
```

### 响应

根据推荐从高到低返回岗位id

```JSON
{
    "job": [6, 1, 2]
}
```

## **3.  根据岗位描述对候选人简历进行打分和筛选**

### **接口描述**

招聘者上传岗位描述，并批量上传候选人的简历，系统根据岗位需求分析简历并打分，返回最合适的候选人。

- **URL**: `/rank_candidates`
- **请求方式**: `POST`

### 请求体

```JSON
{
    "job": {
        "title": "算法工程师",
        "description": "团队介绍：Data-电商团队，负责电商创新项目的算法和大数据工作。依托于字节跳动产品，帮助用户发现并获得好物，享受美好生活。在这个团队，我们不仅要通过推荐和搜索算法帮助用户买到感兴趣的好东西，也要通过风控算法和智能平台治理算法去甄别违规行为，保护用户的购物体验；我们还要建设智能客服技术、大规模商品知识图谱来提升各个交易环节的效率；我们也要结合机器学习和运筹算法，来优化供应链和物流的效率和成本，并进一步提升用户体验；另外我们还会用人工智能来帮助商家提升经营能力。我们的使命：没有难卖的优价好物，让美好生活触手可得。",
        "demand": "1、2025届获得本科及以上学历，计算机相关专业；\n2、扎实的算法和数据结构基础，优秀的编码能力；\n3、机器学习基础扎实，熟悉CF、MF、FM、Word2vec、LR、GBDT、DNN、Wide&Deep等常用的算法模型，熟悉C++/Python/Java等语言，熟悉Linux开发环境；\n4、有个性化推荐、广告、信息检索、自然语言处理、机器学习等相关领域研究或者项目实践经验更佳；\n5、在KDD、NeurIPS、WWW、SIGIR、WSDM、ICML、IJCAI、AAAI、RecSys等会议发表过论文，或者有过数据挖掘/机器学习相关的竞赛经历更佳；\n6、有钻研精神，主观能动性强，能适应快速变化的业务需求，具备良好的团队合作精神和沟通技巧。",
        "location": "上海",
        "company": "比特跳动",
        "salary": "40k*15",
        "job_type": "研发"
    },
    
    "resumes": [
        {
            "id": 1,
            "user_id": 3,
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
        },
        {
            "id": 3,
            "user_id": 4,
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
    ]
}
```

### 响应

返回每个候选人的得分

```JSON
{
    "score": [
        {"id": 3, "score": 80},
        {"id": 4, "score": 70}
    ]
}
```

## **错误处理**

### **常见状态码**

- **200 OK**: 请求成功，返回对应的结果。
- **400 Bad Request**: 请求参数无效，可能是缺少必填字段或文件格式错误。
- **500 Internal Server Error**: 服务器内部错误，通常与文件解析或 AI 服务异常有关。

### **错误响应示例**

```JSON
{
  "error": "Missing required field: resume"
}
```