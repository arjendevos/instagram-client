-- ************************************** instagram_accounts
CREATE TABLE instagram_accounts
(
 id                       serial NOT NULL UNIQUE,
 is_business_account     boolean NOT NULL,
 is_professional_account boolean NOT NULL,
 is_verified             boolean NOT NULL,
 username                varchar(255) NOT NULL UNIQUE,
 full_name               varchar(30) NULL,
 profile_picture         text NOT NULL,
 post_count              int NOT NULL,
 followers_count         bigint NOT NULL,
 following_count         int NOT NULL,
 external_url            varchar(255) NULL,
 business_email          varchar(255) NULL,
 business_phone          varchar(255) NULL,
 business_category       varchar(255) NULL,
 category_name           varchar(255) NULL,
 biography               varchar(255) NULL,
 email                   varchar(255) NULL,
 email_provider          varchar(50) NULL,
 instagram_id            varchar(255) NOT NULL UNIQUE,
 has_management boolean NOT NULL,
 CONSTRAINT PK_5 PRIMARY KEY ( "id" )
);
