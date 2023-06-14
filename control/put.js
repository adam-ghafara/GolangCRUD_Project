var myHeaders = new Headers();
myHeaders.append("Content-Type", "text/plain");

var raw = "{\r\n    \"id_jenis\": 0,\r\n    \"nama_kendaraan\": \"Mazda CX3\",\r\n    \"nomor_kendaraan\": 5543,\r\n    \"detail_servis\": \"Bumper samping tergores, perlu di poles.\",\r\n    \"status_servis\": \"Selesai\"\r\n}";

var requestOptions = {
  method: 'PUT',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("localhost:8000/autobros/:id", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));