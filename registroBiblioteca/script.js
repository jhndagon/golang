var xmlhttp;

function init() {
    xmlhttp = new XMLHttpRequest();

}
function consultar1() {
    var usuario = document.getElementById("usuario1");
    xmlhttp = new XMLHttpRequest();
    var url = "http://localhost:8000/personas/" + usuario.value;
    xmlhttp.open("GET", url, true);
    xmlhttp.send(null);
    xmlhttp.onreadystatechange = function () {

        if (xmlhttp.readyState == 4) {
            if (xmlhttp.status == 200) {
                var det = eval("(" + xmlhttp.responseText + ")");
                if (det["nombres"] != undefined) {
                    document.getElementById("resultadoConsulta").innerHTML = "El usuario ingresado existe.";
                    document.getElementById("nombreU").innerHTML = "Nombre: " + det["nombres"];
                    document.getElementById("usuarioU").innerHTML = "Usuario: " + det["usuario"];
                    document.getElementById("correoU").innerHTML = "Correo: " + det["correo"];
                }
                else {
                    document.getElementById("resultadoConsulta").innerHTML = "El usuario ingresado no existe.";
                }
            }
        }
    };
}
function registrar1() {

    var nombre = document.getElementById("nombres");
    var usuario = document.getElementById("usuario");
    var correo = document.getElementById("correo");
    var contrasena = document.getElementById("contrasena");

    var data = {};
    data.nombres = nombre.value;
    data.correo = correo.value;
    data.usuario = usuario.value;
    data.contrasena = contrasena.value;

    console.log(data);
    xmlhttp = new XMLHttpRequest();
    var url = "http://localhost:8000/personas";
    var json = JSON.stringify(data);
    xmlhttp.open("POST", url, true);
    xmlhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded; charset=utf-8');
    xmlhttp.send(json);
    xmlhttp.onload = function () {
        var users = JSON.parse(xmlhttp.responseText);
        //document.getElementById("resultado").innerHTML = "Los datos se registraron satisfactoriamente.";
        if (xmlhttp.readyState == 4 && xmlhttp.status == "200") {
            document.getElementById("resultado").innerHTML = "Los datos se registraron satisfactoriamente.";
        } else {
        }
    }

}

init()