function getData() {
  owner_name = document.getElementById("owner_name").value;
  vec_type = document.querySelector('input[name="default-radio"]:checked').value;
  vec_name = document.getElementById("vec_name").value;
  vec_code = document.getElementById("vec_code").value;
  vec_detail = document.getElementById("vec_detail").value;
  sendData(owner_name, vec_type, vec_name, vec_code, vec_detail);
}

function sendData(owner_name,vec_type,vec_name,vec_code,vec_detail){
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
    "id_jenis": vec_type,
    "nama_pemilik": owner_name,
    "nama_kendaraan": vec_name,
    "nomor_kendaraan": vec_code,
    "detail_servis": vec_detail,
    "status_servis": "Pending"
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};
  fetch("http://localhost:8000/autobros/tambahservis", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error))
}