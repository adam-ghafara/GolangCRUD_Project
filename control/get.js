var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };
  
  fetch("http://localhost:8000/autobros/tabelservis", requestOptions)
    .then(response => response.text())
    .then(result => collectFunc(result))
    .catch(error => console.log('error', error));

function collectFunc(result) {
  console.log(result);
  var getData=JSON.parse(result);
  postData=getData.forEach(newData);
}

function newData(item) {
  var konten = document.getElementById('konten');

  var row = document.createElement('tr');

  var ownerNameCell = document.createElement('td');
  ownerNameCell.textContent = item.nama_pemilik;

  var vecNameCell = document.createElement('td');
  vecNameCell.textContent = item.nama_kendaraan;

  var vecTypeCell = document.createElement('td');
  vecTypeCell.textContent = item.jenis_kendaraan;

  var vecCodeCell = document.createElement('td');
  vecCodeCell.textContent = item.nomor_kendaraan;

  var dateServisCell = document.createElement('td');
  dateServisCell.textContent = item.tanggal_servis;

  var detailCell = document.createElement('td');
  var detailButton = document.createElement('button');
  detailButton.textContent = 'Lihat';
  detailButton.addEventListener('click', function() {
    // Handle button click here
    console.log('Button clicked for item:', item);
  });
  detailCell.appendChild(detailButton);

  row.appendChild(ownerNameCell);
  row.appendChild(vecNameCell);
  row.appendChild(vecTypeCell);
  row.appendChild(vecCodeCell);
  row.appendChild(dateServisCell);
  row.appendChild(detailCell);

  konten.appendChild(row);
}
