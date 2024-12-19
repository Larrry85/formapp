function addNumbers() {
    // Get user input
    let firstNumber = parseInt(document.getElementById("firstNumber").value);
    let secondNumber = parseInt(document.getElementById("secondNumber").value);

    // Check if inputs are valid numbers
    if (isNaN(firstNumber) || isNaN(secondNumber)) {
        document.getElementById("result").textContent = "Invalid input. Please enter valid numbers.";
    } else {
        // Add numbers
        let sum = firstNumber + secondNumber;
        
        // Display result as a string
        document.getElementById("result").textContent = "The sum of " + firstNumber + " and " + secondNumber + " is: " + sum;
    }
}

document.getElementById("myForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Prevent the form from submitting normally

    var input = document.getElementById("input").value;
    
    // Update the content of the output element with the input value
    document.getElementById("output").innerText = "Input: " + input;

    // Clear the input field after displaying the input
    document.getElementById("input").value = "";
});