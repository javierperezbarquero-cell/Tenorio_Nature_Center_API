CREATE TABLE EmpresaCliente(
    idEmpresaCliente INT AUTO_INCREMENT NOT NULL,
    razonSocial VARCHAR(100) NOT NULL,
    cedulaJuridica VARCHAR(25) NOT NULL,
    telefono VARCHAR(20) DEFAULT NULL,
    email VARCHAR(150) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_empresaCliente PRIMARY KEY(idEmpresaCliente),
    CONSTRAINT uq_empresa_cedulaJuridica UNIQUE(cedulaJuridica)
) ENGINE=INNODB;


CREATE TABLE Cliente(
    idCliente INT AUTO_INCREMENT NOT NULL,
    idEmpresaCliente INT DEFAULT NULL,
    nombre VARCHAR(45) NOT NULL,
    identificador VARCHAR(25) NOT NULL,
    fechaNac DATE NOT NULL,
    telefono VARCHAR(20) NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    fechaRegistro DATE NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_cliente PRIMARY KEY(idCliente),
    CONSTRAINT uq_cliente_identificador UNIQUE(identificador),
    CONSTRAINT fk_cliente_empresa FOREIGN KEY(idEmpresaCliente) REFERENCES EmpresaCliente(idEmpresaCliente)
) ENGINE=INNODB;


CREATE TABLE EmailCliente(
    idEmailCliente INT AUTO_INCREMENT NOT NULL,
    idCliente INT NOT NULL,
    email VARCHAR(150) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_emailCliente PRIMARY KEY(idEmailCliente),
    CONSTRAINT fk_email_cliente FOREIGN KEY(idCliente) REFERENCES Cliente(idCliente)
) ENGINE=INNODB;


CREATE TABLE Guia(
    idGuia INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    identificador VARCHAR(25) NOT NULL,
    fechaNac DATE NOT NULL,
    telefono VARCHAR(20) NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    email VARCHAR(150) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_guia PRIMARY KEY(idGuia),
    CONSTRAINT uq_guia_identificador UNIQUE(identificador)
) ENGINE=INNODB;


CREATE TABLE Idioma(
    idIdioma INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(25) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_idioma PRIMARY KEY(idIdioma),
    CONSTRAINT uq_idioma_nombre UNIQUE(nombre)
) ENGINE=INNODB;


CREATE TABLE IdiomaGuia(
    idIdiomaGuia INT AUTO_INCREMENT NOT NULL,
    idGuia INT NOT NULL,
    idIdioma INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_idiomaGuia PRIMARY KEY(idIdiomaGuia),
    CONSTRAINT uq_idiomaGuia UNIQUE(idGuia, idIdioma),
    CONSTRAINT fk_idiomaGuia_guia FOREIGN KEY(idGuia) REFERENCES Guia(idGuia),
    CONSTRAINT fk_idiomaGuia_idioma FOREIGN KEY(idIdioma) REFERENCES Idioma(idIdioma)
) ENGINE=INNODB;


CREATE TABLE Chofer(
    idChofer INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    identificador VARCHAR(25) NOT NULL,
    fechaNac DATE NOT NULL,
    telefono VARCHAR(20) NOT NULL,
    email VARCHAR(150) NOT NULL,
    tipoLicencia VARCHAR(10) NOT NULL,
    nacionalidad VARCHAR(25) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_chofer PRIMARY KEY(idChofer),
    CONSTRAINT uq_chofer_identificador UNIQUE(identificador)
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


CREATE TABLE Ubicacion(
    idUbicacion INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    direccion VARCHAR(150) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_ubicacion PRIMARY KEY(idUbicacion)
) ENGINE=INNODB;


CREATE TABLE Tour(
    idTour INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    descripcion VARCHAR(150) NOT NULL,
    horario VARCHAR(45) NOT NULL,
    duracion INT NOT NULL,
    cuposMaximos INT NOT NULL,
    precioBase DECIMAL(10,2) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_tour PRIMARY KEY(idTour)
) ENGINE=INNODB;


CREATE TABLE EstadoReserva(
    idEstadoReserva INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(20) NOT NULL,
    descripcion VARCHAR(100) DEFAULT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_estadoReserva PRIMARY KEY(idEstadoReserva),
    CONSTRAINT uq_estadoReserva_nombre UNIQUE(nombre)
) ENGINE=INNODB;


CREATE TABLE Reserva(
    idReserva INT AUTO_INCREMENT NOT NULL,
    idEstadoReserva INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_reserva PRIMARY KEY(idReserva),
    CONSTRAINT fk_reserva_estadoReserva FOREIGN KEY(idEstadoReserva) REFERENCES EstadoReserva(idEstadoReserva)
) ENGINE=INNODB;


CREATE TABLE DetalleReserva(
    idDetalleReserva INT AUTO_INCREMENT NOT NULL,
    idReserva INT NOT NULL,
    idTour INT NOT NULL,
    idGuia INT NOT NULL,
    idChofer INT NOT NULL,
    idUbicacion INT NOT NULL,
    idIdioma INT NOT NULL,
    fechaTour DATE NOT NULL,
    precioUnitario DECIMAL(10,2) NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_detalleReserva PRIMARY KEY(idDetalleReserva),
    CONSTRAINT fk_detalle_reserva FOREIGN KEY(idReserva) REFERENCES Reserva(idReserva),
    CONSTRAINT fk_detalle_tour FOREIGN KEY(idTour) REFERENCES Tour(idTour),
    CONSTRAINT fk_detalle_guia FOREIGN KEY(idGuia) REFERENCES Guia(idGuia),
    CONSTRAINT fk_detalle_chofer FOREIGN KEY(idChofer) REFERENCES Chofer(idChofer),
    CONSTRAINT fk_detalle_ubicacion FOREIGN KEY(idUbicacion) REFERENCES Ubicacion(idUbicacion),
    CONSTRAINT fk_detalle_idioma FOREIGN KEY(idIdioma) REFERENCES Idioma(idIdioma)
) ENGINE=INNODB;


CREATE TABLE Participante(
    idParticipante INT AUTO_INCREMENT NOT NULL,
    idCliente INT NOT NULL,
    idReserva INT NOT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_participante PRIMARY KEY(idParticipante),
    CONSTRAINT uq_participante UNIQUE(idCliente, idReserva),
    CONSTRAINT fk_participante_cliente FOREIGN KEY(idCliente) REFERENCES Cliente(idCliente),
    CONSTRAINT fk_participante_reserva FOREIGN KEY(idReserva) REFERENCES Reserva(idReserva)
) ENGINE=INNODB;


CREATE TABLE EstadoPago(
    idEstadoPago INT AUTO_INCREMENT NOT NULL,
    nombre VARCHAR(20) NOT NULL,
    descripcion VARCHAR(100) DEFAULT NULL,
    fechaCreacion DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_estadoPago PRIMARY KEY(idEstadoPago),
    CONSTRAINT uq_estadoPago_nombre UNIQUE(nombre)
) ENGINE=INNODB;


CREATE TABLE Factura(
    idFactura        INT AUTO_INCREMENT NOT NULL,
    idEstadoPago     INT NOT NULL,
    numeroFactura    VARCHAR(50) NOT NULL,
    fechaFactura     DATE NOT NULL,
    metodoPago       VARCHAR(20) NOT NULL,
    moneda           VARCHAR(30) NOT NULL,
    fechaPago        DATE DEFAULT NULL,
    subtotal         DECIMAL(10,2) NOT NULL,
    descuento        DECIMAL(10,2) NOT NULL,
    impuesto         DECIMAL(10,2) NOT NULL,
    precioTotal      DECIMAL(10,2) NOT NULL,
    fechaCreacion    DATETIME DEFAULT NULL,
    fechaActualizacion DATETIME DEFAULT NULL,
    CONSTRAINT pk_factura PRIMARY KEY(idFactura),
    CONSTRAINT fk_factura_estadoPago FOREIGN KEY(idEstadoPago) REFERENCES EstadoPago(idEstadoPago)
);

CREATE TABLE FacturaParticipante(
    idFacturaParticipante INT AUTO_INCREMENT NOT NULL,
    idFactura             INT NOT NULL,
    idParticipante        INT NOT NULL,
    fechaCreacion         DATETIME DEFAULT NULL,
    fechaActualizacion    DATETIME DEFAULT NULL,
    CONSTRAINT pk_facturaParticipante PRIMARY KEY(idFacturaParticipante),
    CONSTRAINT uq_facturaParticipante UNIQUE(idParticipante),
    CONSTRAINT fk_fp_factura      FOREIGN KEY(idFactura)      REFERENCES Factura(idFactura),
    CONSTRAINT fk_fp_participante FOREIGN KEY(idParticipante) REFERENCES Participante(idParticipante)
);

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