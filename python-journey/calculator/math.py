# function to add
def add(x, y):
    return x + y

# function to subtract
def subtract(x, y):
    return x - y

# function to multiply
def multiply(x, y):
    return x * y

def divide(x, y):
    if y == 0:
        return "Error! Division by zero."
    return x / y

history = []

while True:
    print("\n---Simple Python Calculator---")
    print("1. Add")
    print("2. Subtract")
    print("3. Multiply")
    print("4. Divide")
    print("5. View Calculation History")
    print("6. Quit")

    choice = input("Enter your choice: ").strip()

    if choice.lower() == 'q':
        print("Goodbye!")
        break

    elif choice == '5':
        if not history:
            print("\n [History is empty]")
        else:
            print("\n --- Past Calculations---")
            for record in history:
                print(record)
        continue

    elif choice in ('1', '2', '3', '4',):
        try:
            num1 = float(input("Enter first number: "))
            num2 = float(input("Enter second number: "))
        except ValueError:
            print("Invalid input! please enter numbers only. ")
            continue

        if choice == '1':
            result = add(num1, num2)
            log = f"{num1} + {num2} = {result}"
        elif choice == '2':
            result = subtract(num1, num2)
            log = f"{num1} - {num2} = {result}"
        elif choice == '3':
            result = multiply(num1, num2)
            log = f"{num1} * {num2} = {result}"
        elif choice == '4':
            result = divide(num1, num2)
            log = f"{num1} / {num2} = {result}"

        print(f"Result: {result}")
        history.append(log)
    else:
        print("Invalid Choice. Please choose from the menu.")