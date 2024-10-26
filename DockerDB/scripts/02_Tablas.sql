CONNECT RES/res@//localhost:1521/RESERVASDB;

CREATE TABLE RES.Usuarios (
    id VARCHAR(50) PRIMARY KEY,
    p_nombre VARCHAR(50) NOT NULL,
    s_nombre VARCHAR(50),
    p_apellido VARCHAR(50) NOT NULL,
    s_apellido VARCHAR(50),
    identificacion VARCHAR(15) NOT NULL,
    direccion VARCHAR(100) NOT NULL,
    nacionalidad VARCHAR(30) NOT NULL,
    nick_name VARCHAR(30),
    celular VARCHAR(15) NOT NULL,
    f_nacimiento VARCHAR(10) NOT NULL,
    pais VARCHAR(50) NOT NULL,
    depto VARCHAR(50) NOT NULL,
    email VARCHAR(100)
);

INSERT INTO RES.Usuarios VALUES (124567, 'yar', 'leison', 'palo', 'meque', 987654, 'calle lejana', 'americano', 'rico', 3132224455, '1995-12-01', 'COL', 'ANT', 'ricoelpalo@gmail.com');

