import random
from faker import Faker
from faker.providers import DynamicProvider
import requests

faker = Faker('zh_CN')
base_url = 'https://api.recruitment.kkkstra.cn/api/v1'

job_types=[
    "研发", "销售", "产品", "人力资源", "财务", "运营", "客服", "设计", "项目管理"
]

job_descriptions = [
    "负责技术创新和产品研发，确保产品功能和性能的不断提升。参与需求分析，完成核心代码编写，优化系统架构，支持产品迭代与更新，确保技术方案的可行性和稳定性。",
    "拓展并维护客户关系，负责产品和服务的市场推广与销售。制定并执行销售计划，确保销售目标达成。协助客户解决问题，增加客户满意度，提升公司市场份额。",
    "负责产品规划和设计，分析市场需求，制定产品路线图。与研发、运营等部门合作，确保产品从概念到发布的高效落地，不断提升用户体验和产品竞争力。",
    "负责招聘、培训、绩效考核等工作，确保公司人才资源的合理配置与发展。推动员工职业成长和组织文化建设，支持公司战略目标的实现。",
    "负责公司财务管理和预算控制，确保公司资产的安全和有效利用。制定财务计划，监控财务风险，提供财务数据支持决策，为公司持续稳定发展提供保障。",
    "管理公司业务流程和资源配置，确保公司运营效率和服务质量。通过数据分析优化运营策略，解决运营问题，提升用户体验和公司收益。",
    "负责客户咨询、问题处理和售后支持，确保客户问题及时有效解决。维护客户关系，提高客户满意度，反馈产品和服务问题，协助优化产品和流程。",
    "负责视觉设计和用户界面优化，确保产品的美观性和易用性。参与创意设计，确保品牌风格一致性，支持市场推广和产品宣传。",
    "负责项目的统筹和执行，制定项目计划，分配资源，协调各方，确保项目按时按质完成。监控项目进度，识别并解决风险，确保项目目标达成。"
]

job_demands = [
    "熟悉编程语言和开发工具，具备系统架构和模块设计经验。具备较强的学习能力和解决问题的能力，能够适应快速变化的技术环境，具有良好的团队协作意识。",
    "具备出色的沟通表达能力和客户关系管理经验，熟悉销售流程和谈判技巧。自驱力强，能够承受压力并保持积极心态，有较强的市场敏锐度和开拓精神。",
    "熟悉产品生命周期管理，具备需求分析和市场调研能力，能够清晰地定义产品功能和路线图。具备良好的跨部门沟通能力，能够将客户需求转化为产品方案。",
    "具备招聘、培训、绩效考核等人力资源领域的实际操作经验，熟悉劳动法规和人力资源相关政策。具备较强的沟通能力和亲和力，能够推动组织文化建设。",
    "具备财务分析和预算管理能力，熟悉会计准则和税务政策，能够进行财务报表分析。具有良好的数据敏感度和风险控制意识，能够支持业务决策和资源配置。",
    "具备数据分析和运营策略制定能力，熟悉业务流程管理和资源优化配置。具备良好的项目管理和沟通协调能力，能够推动跨部门合作，提升运营效率。",
    "具备良好的客户服务意识和沟通技巧，熟悉客户关系管理系统，有较强的耐心和应变能力。能够高效处理客户问题，提升客户满意度和忠诚度。",
    "具备出色的审美能力和设计技能，熟悉设计工具和用户体验设计方法。具备品牌理解力，能够创造符合品牌形象的视觉方案，有较强的创新和执行能力。",
    "具备项目管理经验，熟悉项目计划制定、进度监控、风险管理等流程。具备良好的沟通和协调能力，能够有效调动资源，推动项目按计划顺利进行。"
]

resume_descriptions = [
    "拥有五年软件开发经验，精通多种编程语言，专注于系统架构设计和性能优化。具有快速学习和解决问题的能力，乐于接受新挑战并保持对技术的热情。能够在团队中发挥积极作用，推动项目高效完成。",
    "具备丰富的销售经验和扎实的市场推广技巧，善于通过数据分析制定销售策略，达成目标。擅长客户沟通，具备较强的开拓精神和承压能力，乐于在竞争中寻求突破，推动客户满意度和公司业绩双赢。",    
    "在产品管理领域积累了深厚的实战经验，熟悉产品从概念设计到市场发布的全流程。具备敏锐的市场洞察力和用户需求分析能力，善于跨部门协作，致力于为用户提供优质的产品体验和价值。",
    "具备多年人力资源管理经验，熟悉招聘、培训、绩效考核等各项流程，能够精准识别和发展人才。拥有良好的沟通能力和亲和力，擅长推动组织文化建设，致力于提升员工的满意度和公司凝聚力。",
    "拥有扎实的财务管理经验，熟悉预算控制、财务分析和风险管理。具备敏锐的财务数据解读能力和风险控制意识，能够为公司发展提供有力的数据支持和决策建议，确保财务安全和健康发展。"
]

resume_educations = [
    [
        {
            "school": "示例大学",
            "major": "计算机科学",
            "degree": 1,
            "start_time": "2014-09-01T00:00:00Z",
            "end_time": "2018-06-01T00:00:00Z"
        },
        {
            "school": "示例大学",
            "major": "计算机科学",
            "degree": 2,
            "start_time": "2018-09-01T00:00:00Z",
            "end_time": "2020-06-01T00:00:00Z"
        }
    ],
    [
        {
            "school": "国家技术学院",
            "major": "机械工程",
            "degree": 1,
            "start_time": "2013-09-01T00:00:00Z",
            "end_time": "2017-06-01T00:00:00Z"
        },
        {
            "school": "国家技术学院",
            "major": "机械工程",
            "degree": 2,
            "start_time": "2017-09-01T00:00:00Z",
            "end_time": "2019-06-01T00:00:00Z"
        },
        {
            "school": "国家技术学院",
            "major": "机械工程",
            "degree": 3,
            "start_time": "2019-09-01T00:00:00Z",
            "end_time": "2022-06-01T00:00:00Z"
        }
    ],
    [
        {
            "school": "示例商学院",
            "major": "工商管理",
            "degree": 1,
            "start_time": "2012-09-01T00:00:00Z",
            "end_time": "2016-06-01T00:00:00Z"
        },
        {
            "school": "示例商学院",
            "major": "工商管理",
            "degree": 2,
            "start_time": "2016-09-01T00:00:00Z",
            "end_time": "2018-06-01T00:00:00Z"
        }
    ],
    [
        {
            "school": "高级科学研究院",
            "major": "物理学",
            "degree": 1,
            "start_time": "2011-09-01T00:00:00Z",
            "end_time": "2015-06-01T00:00:00Z"
        },
        {
            "school": "高级科学研究院",
            "major": "物理学",
            "degree": 2,
            "start_time": "2015-09-01T00:00:00Z",
            "end_time": "2017-06-01T00:00:00Z"
        },
        {
            "school": "高级科学研究院",
            "major": "物理学",
            "degree": 3,
            "start_time": "2017-09-01T00:00:00Z",
            "end_time": "2020-06-01T00:00:00Z"
        }
    ],
    [
        {
            "school": "艺术设计大学",
            "major": "平面设计",
            "degree": 1,
            "start_time": "2010-09-01T00:00:00Z",
            "end_time": "2014-06-01T00:00:00Z"
        },
        {
            "school": "艺术设计大学",
            "major": "平面设计",
            "degree": 2,
            "start_time": "2014-09-01T00:00:00Z",
            "end_time": "2016-06-01T00:00:00Z"
        }
    ]
]

resume_experiences = [
    [
        {
            "company": "科技创新有限公司",
            "position": "软件开发工程师",
            "start_time": "2019-07-01T00:00:00Z",
            "end_time": "2021-12-31T00:00:00Z"
        },
        {
            "company": "未来科技公司",
            "position": "高级软件工程师",
            "start_time": "2022-01-01T00:00:00Z",
            "end_time": "2023-12-31T00:00:00Z"
        }
    ],
    [
        {
            "company": "卓越市场公司",
            "position": "市场专员",
            "start_time": "2017-05-01T00:00:00Z",
            "end_time": "2019-04-30T00:00:00Z"
        },
        {
            "company": "创新商务集团",
            "position": "市场经理",
            "start_time": "2019-05-01T00:00:00Z",
            "end_time": "2021-12-31T00:00:00Z"
        }
    ],
    [
        {
            "company": "知名产品公司",
            "position": "产品经理助理",
            "start_time": "2015-09-01T00:00:00Z",
            "end_time": "2017-08-31T00:00:00Z"
        },
        {
            "company": "创新产品有限公司",
            "position": "产品经理",
            "start_time": "2017-09-01T00:00:00Z",
            "end_time": "2020-12-31T00:00:00Z"
        }
    ],
    [
        {
            "company": "人力资源管理公司",
            "position": "招聘专员",
            "start_time": "2014-03-01T00:00:00Z",
            "end_time": "2016-06-30T00:00:00Z"
        },
        {
            "company": "卓越人力资源服务公司",
            "position": "人力资源经理",
            "start_time": "2016-07-01T00:00:00Z",
            "end_time": "2019-12-31T00:00:00Z"
        }
    ],
    [
        {
            "company": "全球财务咨询公司",
            "position": "财务分析师",
            "start_time": "2012-10-01T00:00:00Z",
            "end_time": "2015-09-30T00:00:00Z"
        },
        {
            "company": "国际金融服务公司",
            "position": "财务经理",
            "start_time": "2015-10-01T00:00:00Z",
            "end_time": "2018-12-31T00:00:00Z"
        }
    ]
]

resume_projects = [
    [
        {
            "name": "智能推荐系统",
            "description": "开发了一个基于机器学习的推荐系统，用于电商平台的个性化推荐，提高用户转化率。",
            "start_time": "2020-05-01T00:00:00Z",
            "end_time": "2020-10-01T00:00:00Z"
        },
        {
            "name": "大数据分析平台",
            "description": "搭建了一个数据分析平台，用于实时分析用户行为数据，支持业务决策。",
            "start_time": "2021-01-01T00:00:00Z",
            "end_time": "2021-06-01T00:00:00Z"
        }
    ],
    [
        {
            "name": "市场预测模型",
            "description": "设计并实现了一个基于时间序列的市场预测模型，提高了销售预测的准确性。",
            "start_time": "2018-02-01T00:00:00Z",
            "end_time": "2018-07-01T00:00:00Z"
        },
        {
            "name": "客户行为分析",
            "description": "开发了一个数据挖掘项目，分析客户购买行为，帮助优化产品推荐策略。",
            "start_time": "2019-04-01T00:00:00Z",
            "end_time": "2019-09-01T00:00:00Z"
        }
    ],
    [
        {
            "name": "新产品发布计划",
            "description": "领导团队制定新产品发布策略，涵盖市场调研、推广计划和反馈分析。",
            "start_time": "2020-03-01T00:00:00Z",
            "end_time": "2020-08-01T00:00:00Z"
        },
        {
            "name": "产品优化项目",
            "description": "协助产品优化改进，收集用户反馈并提出功能改进建议，提升用户体验。",
            "start_time": "2021-05-01T00:00:00Z",
            "end_time": "2021-11-01T00:00:00Z"
        }
    ],
    [
        {
            "name": "员工培训系统",
            "description": "开发了一个在线培训系统，用于公司内部员工的技能提升和知识管理。",
            "start_time": "2016-07-01T00:00:00Z",
            "end_time": "2017-02-01T00:00:00Z"
        },
        {
            "name": "组织文化建设项目",
            "description": "组织策划多项活动以促进公司组织文化建设，提升员工凝聚力和归属感。",
            "start_time": "2017-05-01T00:00:00Z",
            "end_time": "2017-12-01T00:00:00Z"
        }
    ],
    [
        {
            "name": "财务风险控制系统",
            "description": "设计并实施了财务风险控制系统，监测和预测财务风险，确保公司财务安全。",
            "start_time": "2015-01-01T00:00:00Z",
            "end_time": "2015-09-01T00:00:00Z"
        },
        {
            "name": "预算管理工具开发",
            "description": "开发了一款预算管理工具，提高了财务团队在预算编制和控制上的效率。",
            "start_time": "2016-02-01T00:00:00Z",
            "end_time": "2016-11-01T00:00:00Z"
        }
    ]
]

# register new user
def register(role: int) -> tuple[str, int]:
    user = {
        "username": faker.user_name(),
        "email": faker.email(),
        "password": "123456",
        "role": role,
        "age": faker.random_int(min=18, max=35),
        "degree": faker.random_int(min=1, max=3),
    }
    url = base_url + '/user'
    response = requests.post(url, json=user)

    assert response.status_code == 200, f"Failed to register user: {response.text}, url={url}"
    response = response.json()
    id = response["data"]["id"]
    print(f"registered user, id={id}")
    return user['username'], id

# login
def login(username: str, password: str = "123456") -> str:
    url = base_url + '/session'
    response = requests.post(url, json={"username": username, "password": password})

    assert response.status_code == 200, f"Failed to login: {response.text}, url={url}"
    response = response.json()
    token = response["data"]["token"]
    print(f"login success, username={username}")
    return token

# post new job
def post_job(company:str, token: str):
    min_salary = faker.random_int(min=5, max=30)
    max_salary = faker.random_int(min=min_salary, max=50)
    idx = faker.random_int(min=0, max=8)
    job = {
        "title": faker.job(),
        "description": job_descriptions[idx],
        "demand": job_demands[idx],
        "location": faker.city_name(),
        "company": company,
        "salary": f"{min_salary}k-{max_salary}k",
        "job_type": job_types[idx],
    }
    url = base_url + '/jobs'
    response = requests.post(url, json=job, headers={"Authorization": f"Bearer {token}"})

    assert response.status_code == 200, f"Failed to post job: {response.text}, url={url}"
    response = response.json()
    id = response["data"]["id"]
    print(f"posted job, id={id}")
    return

# post resume
def post_resume(idx: int, token: str):
    resume = {
        "name": faker.name(),
        "gender": faker.random_int(min=1, max=2),
        "phone": faker.phone_number(),
        "email": faker.email(),
        "wechat": faker.user_name(),
        "state": faker.random_int(min=1, max=4),
        "description": resume_descriptions[idx],
        "education": resume_educations[idx],
        "experience": resume_experiences[idx],
        "project": resume_projects[idx],
    }
    url = base_url + '/resumes'
    response = requests.post(url, json=resume, headers={"Authorization": f"Bearer {token}"})
    
    assert response.status_code == 200, f"Failed to post resume: {response.text}, url={url}"
    response = response.json()
    id = response["data"]["id"]
    print(f"posted resume, id={id}")
    return

def post_application(job_id: int, token: str):
    application = {
        "job_id": job_id,
    }
    url = base_url + '/applications'
    response = requests.post(url, json=application, headers={"Authorization": f"Bearer {token}"})
    
    if response.status_code == 200:
        response = response.json()
        id = response["data"]["id"]
        print(f"posted application, id={id}")
    else:
        print(response.text)

# for _ in range(5):
#     username, id = register(1)
#     token = login(username)
#     company = faker.company()
#     for _ in range(10):
#         post_job(company, token)

for i in range(5):
    username, id = register(2)
    token = login(username)
    post_resume(i, token)
    # get all jobs
    url = base_url + '/jobs'
    response = requests.get(url, headers={"Authorization": f"Bearer {token}"})
    response = response.json()
    jobs = response["data"]["jobs"]
    random.shuffle(jobs)
    for i in range(5):
        post_application(jobs[i]["id"], token)