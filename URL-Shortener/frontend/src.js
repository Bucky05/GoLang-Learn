async function handleURLShortener(e) {
    e.preventDefault();
  
    const input = document.getElementById("url");
    const urlValue = input.value;
    const output = document.getElementById("output");
  
    input.value = "";
    output.textContent = "Loading...";
  
    try {
      const res = await fetch("http://127.0.0.1:8000", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url: urlValue }),
      });
  
      if (!res.ok) {
        throw new Error(`Server responded with ${res.status}`);
      }
  
      const data = await res.json();
      output.textContent = data.message || JSON.stringify(data, null, 2); // customize this
    } catch (err) {
      console.error(err);
      output.textContent = "Error: " + err.message;
    }
};
  
document.getElementById("shortenBtn").addEventListener("click", handleURLShortener);