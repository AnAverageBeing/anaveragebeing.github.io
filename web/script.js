function updateTime() {
  const now = new Date();
  const timeOptions = { hour: 'numeric', minute: '2-digit', second: '2-digit', hour12: true };
  const timeString = now.toLocaleString('en-US', timeOptions);
  document.getElementById('clock').textContent = timeString;
}

setInterval(updateTime, 1000); // Update time every second

var tw = new Typewriter('#typing-text', {
  loop: true,
  delay: 30,
  deleteDelay: 1500
});

var messages = [
  'I love computers!',
  'I Use Arch Linux, BTW!',
  'I love anime uwu!',
  'Coding is fun!'
];

function typeMessages() {
  for (var i = 0; i < messages.length; i++) {
    tw.typeString(messages[i])
      .pauseFor(2000)
      .deleteAll()
      .pauseFor(100)
      .start();
  }
}

typeMessages();
