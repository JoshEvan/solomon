# solomon


![Progress](https://progress-bar.dev/100/?title=Backend)
![Progress](https://progress-bar.dev/0/?title=FrontEnd)


Solomon is a simple web app to support product/item/merchandise gallery.

| Current covered features are |
| ------------- |
| 1. search product by any keyword |
| 2. filter product by price range |
| 3. pagination |
| 4. add / update product  |

## Technical Stack
Built with:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) Clean Architecture
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![ElasticSearch](https://img.shields.io/badge/-ElasticSearch-005571?style=for-the-badge&logo=elasticsearch)
NSQ Message Queue


## System Architecture
```mermaid
block-beta
  columns 2
  block:layer1:2
    columns 2
    uct4("Front End Layer"):2
    ProductListPage ProductDetailPage
    ProductAddPage ProductUpdatePage
  end
  block:groupusecase:2
    block:group1a:1
        columns 1
        uct("HTTP API")
        Filter/Search
        Add
        Update
    end
    %%   space:2
    block:group1b:1
        columns 1
        uct1("Event Bus API")
        InsertToSearchEngine
        UpdateToSearchEngine
        space:2
    end
  end

  block:group2:2
    columns 2
    uct2("Infrastructure Layer"):2
    Docker
    %% pg[("PostgreSQL")] Redis[("Redis")]:1
    PostgreSQL Redis
    ElasticSearch:1
    MQ["Event Bus"]:1
    MonitoringTool
end

class uct BT
class uct1 BT
class uct2 BT
class uct3 BT
class uct4 BT
classDef BT stroke:transparent,fill:transparent,font:center
```

## Code Architecture
There are no backward arrows from inner to outer layer
![image](https://github.com/user-attachments/assets/347bed5d-fb41-4aea-a6f3-0a3da262e63d)
