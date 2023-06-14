function uievent() {
    document.getElementById("dashboard").addEventListener("click", function() {
        window.location.href = "/autobros/dashboard";
      });
    document.getElementById("tambahservis").addEventListener("click", function() {
        window.location.href = "/autobros/tambahservis";
      });
    document.getElementById("tabelservis").addEventListener("click", function() {
        window.location.href = "/autobros/tabelservis";
    });
}