CREATE TABLE mahasiswa (
                           id_mahasiswa integer PRIMARY KEY,
                           nama_mahasiswa varchar(255)
);

CREATE TABLE matakuliah (
                            id_matakuliah integer PRIMARY KEY,
                            nama_matakuliah varchar(255)
);

CREATE TABLE nilai (
                       id_nilai integer,
                       id_mahasiswa integer,
                       id_matakuliah integer,
                       nilai integer,
                       CONSTRAINT fk_mahasiswa
                           FOREIGN KEY(id_mahasiswa)
                               REFERENCES mahasiswa(id_mahasiswa)
                               ON DELETE CASCADE,
                       CONSTRAINT fk_matakuliah
                           FOREIGN KEY(id_matakuliah)
                               REFERENCES matakuliah(id_matakuliah)
                               ON DELETE CASCADE
);

INSERT INTO
    mahasiswa (id_mahasiswa, nama_mahasiswa)
VALUES
    (1001,'Budi'),
    (1002,'Aris'),
    (1003,'Panji');

INSERT INTO
    matakuliah (id_matakuliah, nama_matakuliah)
VALUES
    (101,'Struktur Data'),
    (102,'Rangkaian Digital'),
    (103,'Aljabar linear');

INSERT INTO
    nilai (id_nilai, id_mahasiswa, id_matakuliah, nilai)
VALUES
    (1000001,1001, 101, 85),
    (1000002,1001, 102, 75),
    (1000003,1001, 103, 70),
    (1000004,1002, 101, 79),
    (1000005,1002, 102, 55),
    (1000006,1002, 103, 90),
    (1000007,1003, 101, 73),
    (1000008,1003, 102, 81),
    (1000009,1003, 103, 61);
