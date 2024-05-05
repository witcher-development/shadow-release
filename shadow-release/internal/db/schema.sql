create table if not exists app(
	id INTEGER PRIMARY KEY
);

create table if not exists version(
	id integer PRIMARY KEY,
	name text not null,
	app integer not null,
	FOREIGN KEY(app) REFERENCES app(id)
);

create table if not exists record(
	id integer PRIMARY KEY,
	version integer not null,
	path text not null,
	method text not null,
	reqbody text not null,
	resbody text not null,
	synckey text not null,
	created_at datetime default CURRENT_TIMESTAMP,
	FOREIGN KEY(version) REFERENCES version(id)
);
