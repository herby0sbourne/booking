// class Notification {
//   constructor() {
//     this.toast = function () {
//       console.log('click button and got post');
//     };

//     this.success = () => {
//       console.log('success');
//     };
//   }

//   hey() {
//     console.log('hello');
//   }
// }

let attention = Prompt();
// let note = new Notification();

(function () {
  "use strict";
  window.addEventListener(
    "load",
    function () {
      // Fetch all the forms we want to apply custom Bootstrap validation styles to
      let forms = document.getElementsByClassName("needs-validation");
      // Loop over them and prevent submission
      Array.prototype.filter.call(forms, function (form) {
        form.addEventListener(
          "submit",
          function (event) {
            if (form.checkValidity() === false) {
              event.preventDefault();
              event.stopPropagation();
            }
            form.classList.add("was-validated");
          },
          false
        );
      });
    },
    false
  );
})();

function notify(msg, msgType) {
  notie.alert({
    type: msgType,
    text: msg,
    // stay:true
  });
}

function modal(title, text, icon, confirmButtonText) {
  Swal.fire({
    icon: icon,
    title: title,
    html: text,
    footer: '<a href="">Why do I have this issue?</a>',
    confirmButtonText: confirmButtonText,
  });
}

function Prompt() {
  let toast = (options) => {
    const { msg = "", icon = "success", position = "top-end" } = options;
    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      position,
      icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({});
  };

  let status = (options) => {
    const { title = "", msg = "", icon = "success" } = options;

    Swal.fire({
      icon: icon === "success" ? "success" : "error",
      title,
      text: msg,
    });
  };

  async function custom(options) {
    const { html = "", title = "" } = options;

    const { value: result } = await Swal.fire({
      title,
      html,
      focusConfirm: false,
      backdrop: false,
      showCancelButton: true,
      inputAutoFocus: false,
      willOpen: () => {
        if (options.willOpen === undefined) return;
        options.willOpen();
      },
      preConfirm: () => {
        return [
          document.getElementById("start-modal").value,
          document.getElementById("end-modal").value,
        ];
      },
      didOpen: () => {
        if (options.didOpen === undefined) return;
        options.didOpen();
      },
    });
    // console.log(result);
    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result[0] === "" || result[1] === "") return;

        if (options.callback !== undefined) {
          options.callback(result);
        } else {
          options.callback(false);
        }
      } else {
        options.callback(false);
      }
    }

    // if (formValues) {
    //     Swal.fire(JSON.stringify(formValues));
    // }
  }

  return {
    toast,
    status,
    custom,
  };
}
