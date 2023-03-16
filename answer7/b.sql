SELECT nama_mahasiswa, nama_matakuliah, nilai,
       Case
           WHEN nilai >70 THEN 'lulus'
           ELSE 'tidak lulus'
           END ket_lulus
FROM nilai
         INNER JOIN mahasiswa on mahasiswa.id_mahasiswa = nilai.id_mahasiswa
         INNER JOIN matakuliah on matakuliah.id_matakuliah = nilai.id_matakuliah;