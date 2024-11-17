document.addEventListener("DOMContentLoaded", () => {
    const jobForm = document.getElementById("jobForm");
    const addVisitButton = document.getElementById("addVisit");
    const visitsContainer = document.getElementById("visitsContainer");
    const statusForm = document.getElementById("statusForm");
    const statusResult = document.getElementById("statusResult");

    let visitCount = 0;

    addVisitButton.addEventListener("click", () => {
        visitCount++;
        const visitDiv = document.createElement("div");
        visitDiv.className = "mb-3";
        visitDiv.innerHTML = `
            <label for="visit-${visitCount}" class="form-label">Visit ${visitCount}</label>
            <div class="input-group mb-2">
                <input type="text" class="form-control" name="store_id" placeholder="Store ID" required>
                <input type="text" class="form-control" name="image_url" placeholder="Image URL (comma-separated)" required>
                <input type="text" class="form-control" name="visit_time" placeholder="Visit Time" required>
            </div>
        `;
        visitsContainer.appendChild(visitDiv);
    });

    jobForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        const count = document.getElementById("count").value;
        const visits = [];
        Array.from(visitsContainer.children).forEach((child) => {
            const inputs = child.querySelectorAll("input");
            const storeId = inputs[0].value;
            const imageUrls = inputs[1].value.split(",");
            const visitTime = inputs[2].value;

            visits.push({ store_id: storeId, image_url: imageUrls, visit_time: visitTime });
        });

        const payload = { count, visits };

        try {
            const response = await fetch("/api/submit/", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload),
            });

            if (response.status === 201) {
                const data = await response.json();
                alert(`Job created successfully! Job ID: ${data.job_id}`);
            } else {
                alert("Error creating job. Please check the inputs.");
            }
        } catch (error) {
            console.error("Error:", error);
            alert("An error occurred. Please try again.");
        }
    });

    statusForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        const jobId = document.getElementById("jobId").value;

        try {
            const response = await fetch(`/api/status?jobid=${jobId}`, { method: "GET" });

            if (response.status === 200) {
                const data = await response.json();
                statusResult.innerHTML = `<p>Status: ${data.status}</p>`;
            } else {
                alert("Job ID not found or invalid.");
            }
        } catch (error) {
            console.error("Error:", error);
            alert("An error occurred. Please try again.");
        }
    });
});
