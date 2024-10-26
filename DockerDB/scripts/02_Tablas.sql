CONNECT RES/res@//localhost:1521/RESERVASDB;

CREATE TABLE RES.Usuariosr(
    id NUMBER(20) PRIMARY KEY,
    p_nombre VARCHAR(50) NOT NULL,
    s_nombre VARCHAR(50),
    p_apellido VARCHAR(50) NOT NULL,
    s_apellido VARCHAR(50),
    nid NUMBER(15) NOT NULL,
    direccion VARCHAR(100) NOT NULL,
    nacionalidad VARCHAR(30) NOT NULL,
    nick_name VARCHAR(30),
    celular NUMBER(15) NOT NULL,
    f_nacimiento DATE NOT NULL,
    pais VARCHAR(50) NOT NULL,
    depto VARCHAR(50) NOT NULL,
    email VARCHAR2(100)

);

INSERT INTO RES.Usuariosr(id, p_nombre, s_nombre, p_apellido, s_apellido, nid, direccion, nacionalidad, nick_name, celular, f_nacimiento, pais, depto, email) 
VALUES (124567, 'yar', 'leison', 'palo', 'meque', 987654, 'calle lejana', 'americano', 'rico', 3132224455, TO_DATE('1995-12-01', 'YYYY-MM-DD'), 'COL', 'ANT', 'ricoelpalo@gmail.com');
