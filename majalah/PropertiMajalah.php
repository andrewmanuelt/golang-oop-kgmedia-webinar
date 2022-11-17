<?php 

require './PropertiMajalahInterface.php';
require './AbstractMajalah.php';

class PropertiMajalah extends AbstractMajalah implements PropertiMajalahInterface {
    public $nama_majalah;
    public $brand;
    public $edisi;
    public $halaman;
    
    public function __construct()
    {
        $this->nama_majalah = '';
        $this->brand = '';
        $this->edisi = '';
        $this->halaman = '';
    }

    public function get_penulis() : string {
        return $this->penulis;
    }

    public function set_penulis($penulis = '') {
        $this->penulis = $penulis;
    }
    
    public function get_nama_majalah() : string {
        return $this->nama_majalah;
    }

    public function set_nama_majalah($nama_majalah = '') {
        $this->nama_majalah = $nama_majalah;
    }

    public function get_brand() : string {
        return $this->brand;
    }

    public function set_brand($brand = '') {
        $this->brand = $brand;
    }

    public function get_edisi() : string {
        return $this->edisi;
    }

    public function set_edisi($edisi = '') {
        $this->edisi = $edisi;
    }

    public function get_halaman() : string {
        return $this->halaman;
    }
    
    public function set_halaman($halaman = '') {
        $this->halaman = $halaman;
    }
}