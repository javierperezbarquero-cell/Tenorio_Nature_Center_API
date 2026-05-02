CREATE TABLE Cliente(
    idCliente INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    telefono INT NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    fechaRegistro DATE NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_cliente PRIMARY KEY(idCliente)
) ENGINE=INNODB;


CREATE TABLE Guia(
    idGuia INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    fechaNac DATE NOT NULL,
    telefono INT NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    email VARCHAR(45) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_guia PRIMARY KEY(idGuia)
) ENGINE=INNODB;


CREATE TABLE Chofer(
    idChofer INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    fechaNac DATE NOT NULL,
    telefono INT NOT NULL,
    email VARCHAR(45) NOT NULL,
    tipoLicencia VARCHAR(10) NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_chofer PRIMARY KEY(idChofer)
) ENGINE=INNODB;


CREATE TABLE Ubicacion(
    idUbicacion INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    direccion VARCHAR(45) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_ubicacion PRIMARY KEY(idUbicacion)
) ENGINE=INNODB;


CREATE TABLE Idioma(
    idIdioma INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(25) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_idioma PRIMARY KEY(idIdioma),
    CONSTRAINT uq_idioma_nombre UNIQUE(nombre)
) ENGINE=INNODB;


CREATE TABLE Tour(
    idTour INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    descripcion VARCHAR(90) NOT NULL,
    horario VARCHAR(45) NOT NULL,
    duracion INT NOT NULL,
    cuposMaximos INT NOT NULL,
    precioBase DOUBLE NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_tour PRIMARY KEY(idTour)
) ENGINE=INNODB;


CREATE TABLE Transporte(
    idTransporte INT AUTO_INCREMENT NOT NULL,
    idChofer INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_transporte PRIMARY KEY(idTransporte),
    CONSTRAINT fk_transporte_chofer FOREIGN KEY(idChofer) REFERENCES Chofer(idChofer)
) ENGINE=INNODB;


CREATE TABLE Vehiculo(
    idVehiculo INT AUTO_INCREMENT NOT NULL,
    idChofer INT NOT NULL,
    matricula VARCHAR(20) NOT NULL,
    capacidad INT NOT NULL,
    modelo VARCHAR(25) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_vehiculo PRIMARY KEY(idVehiculo),
    CONSTRAINT fk_vehiculo_chofer FOREIGN KEY(idChofer) REFERENCES Chofer(idChofer)
) ENGINE=INNODB;


CREATE TABLE EmailCliente(
    idEmailCliente INT AUTO_INCREMENT NOT NULL,
    email VARCHAR(45) NOT NULL,
    idCliente INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_cliente_email PRIMARY KEY(idEmailCliente),
    CONSTRAINT fk_email_cliente FOREIGN KEY(idCliente) REFERENCES Cliente(idCliente)
) ENGINE=INNODB;


CREATE TABLE IdiomaGuia(
    idIdiomaGuia INT AUTO_INCREMENT NOT NULL,
    idGuia INT NOT NULL,
    idIdioma INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_idioma_guia PRIMARY KEY(idIdiomaGuia),
    CONSTRAINT fk_idioma_guia FOREIGN KEY(idGuia) REFERENCES Guia(idGuia),
    CONSTRAINT fk_idioma_idioma FOREIGN KEY(idIdioma) REFERENCES Idioma(idIdioma)
) ENGINE=INNODB;


CREATE TABLE Reserva(
    idReserva INT AUTO_INCREMENT NOT NULL,
    idCliente INT NOT NULL,
    idTour INT NOT NULL,
    idGuia INT NOT NULL,
    idTransporte INT NOT NULL,
    idUbicacion INT NOT NULL,
    idIdioma INT NOT NULL,
    cantidadPersonas INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_reserva PRIMARY KEY(idReserva),
    CONSTRAINT fk_reserva_cliente FOREIGN KEY(idCliente) REFERENCES Cliente(idCliente),
    CONSTRAINT fk_reserva_tour FOREIGN KEY(idTour) REFERENCES Tour(idTour),
    CONSTRAINT fk_reserva_guia FOREIGN KEY(idGuia) REFERENCES Guia(idGuia),
    CONSTRAINT fk_reserva_transporte FOREIGN KEY(idTransporte) REFERENCES Transporte(idTransporte),
    CONSTRAINT fk_reserva_ubicacion FOREIGN KEY(idUbicacion) REFERENCES Ubicacion(idUbicacion),
    CONSTRAINT fk_reserva_idioma FOREIGN KEY(idIdioma) REFERENCES Idioma(idIdioma)
) ENGINE=INNODB;


CREATE TABLE Participante(
    idParticipante INT AUTO_INCREMENT NOT NULL,
    idReserva INT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    fechaNac DATE NOT NULL,
    nacionalidad VARCHAR(45) NOT NULL,
    telefono INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_participante PRIMARY KEY(idParticipante),
    CONSTRAINT fk_participante_reserva FOREIGN KEY(idReserva) REFERENCES Reserva(idReserva)
) ENGINE=INNODB;


CREATE TABLE Factura(
    idFactura INT AUTO_INCREMENT NOT NULL,
    idReserva INT NOT NULL,
    numeroFactura VARCHAR(25) NOT NULL,
    fechaFactura DATE NOT NULL,
    metodoPago VARCHAR(15) NOT NULL,
    moneda VARCHAR(30) NOT NULL,
    estadoPago VARCHAR(25) NOT NULL,
    fechaPago DATE DEFAULT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    impuesto DECIMAL(10,2) NOT NULL,
    descuento DECIMAL(10,2) NOT NULL,
    precioTotal DECIMAL(10,2) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_factura PRIMARY KEY(idFactura),
    CONSTRAINT uq_factura_reserva UNIQUE(idReserva),
    CONSTRAINT fk_factura_reserva FOREIGN KEY(idReserva) REFERENCES Reserva(idReserva)
) ENGINE=INNODB;


CREATE TABLE usuarios(
    idUsuario INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(150),
    rol VARCHAR(30),
    correo VARCHAR(150) NOT NULL,
    contrasena VARCHAR(255) NOT NULL,
    descripcion TEXT,
    imagen VARCHAR(255),
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    tokenRecordar VARCHAR(255),
    CONSTRAINT pk_usuarios PRIMARY KEY(idUsuario),
    CONSTRAINT uq_usuarios_correo UNIQUE(correo)
) ENGINE=INNODB;