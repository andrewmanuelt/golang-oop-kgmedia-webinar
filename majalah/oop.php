<?php 

include './Majalah.php';

$majalah = new Majalah();

$majalah->set_nama_majalah('Bobo Kids');
$majalah->set_brand('Bobo');
$majalah->set_edisi('20 2022');
$majalah->set_halaman('17');
$majalah->set_penulis('Abdul Kadir');

echo 'Nama majalah : '.$majalah->get_nama_majalah();
echo '<br>';
echo 'Brand majalah : '.$majalah->get_brand();
echo '<br>';
echo 'Edisi majalah : '.$majalah->get_edisi();
echo '<br>';
echo 'Halaman majalah : '.$majalah->get_halaman();
echo '<br>';
echo 'Penulis majalah : '.$majalah->get_penulis();
?>