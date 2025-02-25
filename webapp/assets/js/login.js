$('#login').on('submit', fazerLogin)

function fazerLogin(evento) {
    evento.preventDefault()
    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        window.location = "/home";
        alert("Login realizado com sucesso");
    }).fail(function(erro) {
        console.log(erro)
        alert("Erro ao realizar login");
    });
}