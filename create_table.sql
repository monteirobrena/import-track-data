CREATE TABLE tracks (
  id bigserial NOT NULL,
  name character varying(255) NOT NULL,
  isrc character varying(255) NOT NULL,
  units integer,
  revenue numeric(9),
  PRIMARY KEY (id)
);
