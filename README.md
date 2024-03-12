# Products CRUD

### Software Design
**DDD (Domain-Driven Design)** 

By maintianing an onion structure with domain-application-infrastructure. 

### Data Source
CockroachDB
Redis 

### APIs
* GET `/products`: Retrieve all products.
* GET `/products/:id`: Retrieve a specific product.
* POST `/products`: Create a new product.
* PUT `/products/:id`: Update an existing product.
* DELETE `/products/:id`: Delete an existing product.
* GET `/products/search?name=**` Return a list of products matching the text string.

## Using the Multi-Strategy Logger 

- How to use Logger 
1. Initialise NewLoggerRepository and store inside Persistence 

```	
func NewPersistence() (*Persistence, error) {

	// Initialise others 

    // Initialise Logger
    logger := logger.NewLoggerRepositories([]string{"Honeycomb", "Zap"})

	return &Persistence{
        // Other databases and engines
		Logger:             logger,
	}

}
```

2. Then, call it from the Persistence as `r.p.Logger`

```
func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/GetOrder", map[string]interface{}{"id": id})
	defer span.End()

    var order *entity.Order
	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error
	if err != nil {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	r.p.Logger.Info("get order", map[string]interface{}{"data": order})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, errors.New("order not found")
	}

	return order, nil
}
```

-- Explain Trace 

The overall representation of a request's journey through a system. A trace is composed of one or more spans.

-- Explain Span

A single unit of work within a trace. Each span represents a specific function or operation and contains metadata such as start time, end time, and attributes.

