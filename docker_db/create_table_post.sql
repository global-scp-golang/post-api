CREATE TABLE IF NOT EXISTS public."Posts"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    title character varying(100) COLLATE pg_catalog."default" NOT NULL,
    content character varying(100) COLLATE pg_catalog."default" NOT NULL,
    "createdDt" date NOT NULL,
    "modifiedDt" date NOT NULL,
    CONSTRAINT "Posts_pkey" PRIMARY KEY (id)
)
