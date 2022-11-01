
# Storage

- Có chức năng tương tự như Repository, tương tác với DB.
- Mỗi Action với DB có 1 file riêng, VD create:
```go
package restaurantstorage

import (
	"context"
	"bds-service/common"
	"bds-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

```



```go
package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
```

- return DB connection




# Business

- Gồm các file riêng biệt, mỗi file gồm:
  1. Interface của storage 
    ```go
    type CreateRestaurantStore interface {
    Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
    }
    ```
  2. "Constructor" để tạo biz
    ```go
    func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
        return &createRestaurantBiz{store: store}
    }
    ```
  3. Struct chứa storage, các dependencies cần thiết cho business
    ```go
    type createRestaurantBiz struct {
        store CreateRestaurantStore
    }
    ```
  4. Method cho struct trên: validate data, gọi xuống store
    ```go
    func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
    
        if err := data.Validate(); err != nil {
            return err
        }
    
        err := biz.store.Create(ctx, data)
        return err
    }
    ```






# Transportation

## Nhiệm vụ

```go
func CreateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)

		}
		store := restaurantstorage.NewSQLStore(ctx.GetMainDbConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		err := biz.CreateRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)

		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
```

- Là API handler, nhận AppContext
    1. Parse data
    2. Create storage
    3. Create business, pass storage vào
    4. Gọi business function, trả về error (nếu có)
    5. Response.




