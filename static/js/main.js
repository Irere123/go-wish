tailwind.config = {
  theme: {
    fontFamily: {
      sans: ["Inter", "Poppins"],
    },
  },
};

const button = document.getElementById("create-btn");
const emojis = ["🎉", "👏", "🥳", "🎊", "🎈", "🎇"];

if (button) {
  button.addEventListener("click", () => {
    createWithRequest();
  });
}

function createEmoji() {
  const emoji = document.createElement("div");
  emoji.textContent = emojis[Math.floor(Math.random() * emojis.length)];
  emoji.classList.add("emoji");

  // Random position near the top-left corner
  const x = Math.random() * window.innerWidth * 0.2;
  const y = Math.random() * window.innerHeight * 0.2;

  emoji.style.left = `${x}px`;
  emoji.style.top = `${y}px`;

  // Random directions, size, and rotation
  const randomX = (Math.random() - 0.5) * 2000 + "px";
  const randomY = (Math.random() - 0.5) * 2000 + "px";
  const randomSize = Math.random() * 1.5 + 0.5; // random size between 0.5 and 2
  const randomRotation = Math.random() * 720 + "deg"; // random rotation

  emoji.style.setProperty("--random-x", randomX);
  emoji.style.setProperty("--random-y", randomY);
  emoji.style.setProperty("--random-rotation", randomRotation);
  emoji.style.transform = `scale(${randomSize})`;

  document.body.appendChild(emoji);

  // Remove emoji after animation ends
  setTimeout(() => {
    emoji.remove();
  }, 2000);
}

function createBurst() {
  const burst = document.createElement("div");
  burst.classList.add("burst");
  document.body.appendChild(burst);

  // Remove burst after animation ends
  setTimeout(() => {
    burst.remove();
  }, 800);
}

async function createWithRequest() {
  const name = document.getElementById("name").value;
  const age = Number(document.getElementById("age").value);
  const message = document.getElementById("message").value;

  if (!message || !age || !name) {
    alert("Provide all values");
    return;
  }

  if (typeof age !== "number") {
    alert("Invalid age");
    return;
  }

  createBurst();
  for (let i = 0; i < 50; i++) {
    createEmoji();
  }

  setTimeout(async () => {
    const resp = await fetch("/create-wish", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, age, message }),
    });

    const data = await resp.json();

    if (data) {
      window.location.href = `/success?id=${data.ID}`;
    }
  }, 1000);
}

function fallbackCopyTextToClipboard(text) {
  var textArea = document.createElement("textarea");
  textArea.value = text;

  // Avoid scrolling to bottom
  textArea.style.top = "0";
  textArea.style.left = "0";
  textArea.style.position = "fixed";

  document.body.appendChild(textArea);
  textArea.focus();
  textArea.select();

  try {
    var successful = document.execCommand("copy");
    var msg = successful ? "successful" : "unsuccessful";
    console.log("Fallback: Copying text command was " + msg);
  } catch (err) {
    console.error("Fallback: Oops, unable to copy", err);
  }

  document.body.removeChild(textArea);
}

function copyTextToClipboard(text) {
  if (!navigator.clipboard) {
    fallbackCopyTextToClipboard(text);
    return;
  }

  navigator.clipboard.writeText(text).then(
    function () {
      console.log("Async: Copying to clipboard was successful!");
    },
    function (err) {
      console.error("Async: Could not copy text: ", err);
    }
  );
}

document.getElementById("copyBtn").addEventListener("click", () => {
  const text = document.getElementById("link").innerText;
  copyTextToClipboard(text);
});
