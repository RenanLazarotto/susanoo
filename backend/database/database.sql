CREATE DATABASE job_manager;
USE job_manager;

CREATE TABLE empresas (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	nome VARCHAR(255) NOT NULL UNIQUE,
	cnpj VARCHAR(20) NOT NULL UNIQUE,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE enderecos (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	logradouro VARCHAR(255) NOT NULL UNIQUE,
	numero VARCHAR(10) NOT NULL,
	complemento VARCHAR(100),
	bairro VARCHAR(100) NOT NULL,
	cidade VARCHAR(100) NOT NULL,
	cep VARCHAR(9) NOT NULL,
	estado VARCHAR(20),
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE endereco_empresa (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_empresa INTEGER NOT NULL,
	id_endereco INTEGER NOT NULL,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE contato_empresa (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_empresa INTEGER NOT NULL,
	tipo ENUM("telefone", "whatsapp", "email") NOT NULL,
	contato VARCHAR(255) UNIQUE,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE empregos (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_empresa INTEGER NOT NULL,
	ocupacao_inicial VARCHAR(255) NOT NULL,
	remuneracao_inicial DECIMAL(15,2) NOT NULL,
	tipo_contrato VARCHAR(255) NOT NULL,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE ocupacoes (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_emprego INTEGER NOT NULL,
	ocupacao VARCHAR(255) NOT NULL,
	remuneracao_inicial DECIMAL(15,2) NOT NULL,
	data_inicio DATETIME NOT NULL,
	data_fim DATETIME,
	carga_horaria INTEGER NOT NULL COMMENT "Carga horária definida em minutos",
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE remuneracoes (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_emprego INTEGER NOT NULL,
	id_ocupacao INTEGER NOT NULL,
	remuneracao DECIMAL(15,2) NOT NULL,
	data DATETIME NOT NULL,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE cartao_ponto (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_emprego INTEGER NOT NULL,
	horario DATETIME NOT NULL,
	tipo ENUM("entrada", "saida") NOT NULL,
	saldo INTEGER NOT NULL COMMENT "saldo do dia, pode ser negativo",
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE banco_horas (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_emprego INTEGER NOT NULL,
	data DATETIME NOT NULL,
	saldo INTEGER NOT NULL,
	criado DATETIME NOT NULL,
	atualizado DATETIME,
	apagado DATETIME,
	PRIMARY KEY(id)
);


CREATE TABLE holerites (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_emprego INTEGER NOT NULL,
	id_remuneracao INTEGER NOT NULL,
	referencia DATETIME NOT NULL,
	PRIMARY KEY(id)
);


CREATE TABLE detalhamento_holerite (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	id_holerite INTEGER NOT NULL,
	tipo ENUM("credito", "debito") NOT NULL,
	valor DECIMAL(15,2) NOT NULL,
	descricao TEXT(65535) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE historico (
	id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
	tabela ENUM("banco_horas", "cartao_ponto", "contato_empresa", "detalhamento_holerite", "empregos", "empresas", "enderecos", "endereco_empresa", "holerites", "ocupacoes", "remuneracoes") NOT NULL,
	acao ENUM("INSERT", "UPDATE", "DELETE") NOT NULL,
	descricao TEXT(65535) NOT NULL,
	dados_antigos JSON NOT NULL,
	PRIMARY KEY(id)
);


ALTER TABLE endereco_empresa
ADD FOREIGN KEY(id_empresa) REFERENCES empresas(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE endereco_empresa
ADD FOREIGN KEY(id_endereco) REFERENCES enderecos(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE contato_empresa
ADD FOREIGN KEY(id_empresa) REFERENCES empresas(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE empregos
ADD FOREIGN KEY(id_empresa) REFERENCES empresas(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ocupacoes
ADD FOREIGN KEY(id_emprego) REFERENCES empregos(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE remuneracoes
ADD FOREIGN KEY(id_emprego) REFERENCES empregos(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE remuneracoes
ADD FOREIGN KEY(id_ocupacao) REFERENCES ocupacoes(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE cartao_ponto
ADD FOREIGN KEY(id_emprego) REFERENCES empregos(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE banco_horas
ADD FOREIGN KEY(id_emprego) REFERENCES empregos(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE holerites
ADD FOREIGN KEY(id_remuneracao) REFERENCES remuneracoes(id)
ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE detalhamento_holerite
ADD FOREIGN KEY(id_holerite) REFERENCES holerites(id)
ON UPDATE CASCADE ON DELETE CASCADE;