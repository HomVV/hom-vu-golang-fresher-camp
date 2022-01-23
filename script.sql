CREATE TABLE `food_delivery`.`FoodCategories`(
	`Id` 			int 			auto_increment
	,`Name`	 		nvarchar(50) 	NOT NULL
	,`Description` 	nvarchar(100)
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
);

CREATE TABLE `food_delivery`.`Restaurants`(
	 `Id` 			int 			auto_increment
	,`Name`			nvarchar(50)
	,`Address`		nvarchar(100) 	not NULL
	,`Image` 		json
	,`Verified` 	BIT
	,`TotalRate` 	decimal(10,2)
	,`RateCount` 	int
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
);



CREATE TABLE `food_delivery`.`Food`(
	 `Id` 			int 			auto_increment
	,`Name`	 		nvarchar(50) 	NOT NULL
	,`CategoryId` 	int 			not null
	,`RestaurantId` int 			not null
	,`Title` 		nvarchar(100)
	,`Image` 		Json
	,`Price` 		DECIMAL(10,2)
	,`Quantity` 	int
	,`TotalRate` 	decimal(10,2)
	,`RateCount` 	int
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
	,FOREIGN KEY (`CategoryId`) REFERENCES `FoodCategories`(`Id`)
	,FOREIGN KEY (`RestaurantId`) REFERENCES `Restaurants`(`Id`)
);


CREATE TABLE `food_delivery`.`Customers`(
	`Id` 			int 			auto_increment
	,`Name` 		nvarchar(50) 	not null
	,`Address` 		nvarchar(100) 	not null
	,`PhoneNumber` 	varchar(12) 	not null
	,`Password`		varchar(100)	not null
	,`Avatar` 		json
	,`Email` 		varchar(100)
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
);

CREATE TABLE `food_delivery`.`FoodReviews`(
	 `Id` 			int 			auto_increment
	,`FoodId` 		int
	,`CustomerId` 	int
	,`Content` 		nvarchar(500)
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
	,FOREIGN KEY (`FoodId`) REFERENCES `Food`(`Id`)
	,FOREIGN KEY (`CustomerId`) REFERENCES `Customers`(`Id`)
);

CREATE TABLE `food_delivery`.`RestaurantReviews`(
	 `Id` 			int 			auto_increment
	,`RestaurantId` int
	,`CustomerId` 	int
	,`Content` 		nvarchar(500)
	,`Status` 		Tinyint(3) 		DEFAULT 1
	,`Created_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP
	,`Updated_At` 	TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
	,PRIMARY KEY (`Id`)
	,FOREIGN KEY (`RestaurantId`) REFERENCES `Restaurants`(`Id`)
	,FOREIGN KEY (`CustomerId`) REFERENCES `Customers`(`Id`)
);
