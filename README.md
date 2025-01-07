
# **Excel Formatter GUI App**  
## **How to Install and Run:**  

### **Clone the Repo:**  
```sh
git clone https://github.com/yourusername/excel-formatter-go.git
cd excel-formatter-go
```

### **Run the App:**  
```sh
go run .
```

---

## **How to Use It:**  

1. **"Upload File"**  
   - Click the **"Upload File"** button.  
   - Select an `.xlsx` file from your system.  

2. **"Enter a file name"**  
   - Enter the name of an `.xlsx` file located in your **Downloads folder**.  
   - Example: `example.xlsx`  

---

## **What Happens Next:**  

- **If the file is already formatted:**  
   A dialog box will inform you that the file is already formatted.  

- **If an error occurs:**  
   A dialog box will display the error message.  

- **If the file is successfully formatted:**  
   A dialog box will confirm success, and a new sheet named **"Formatted"** will appear in your Excel file.  

---

## **Dependencies:**  
Ensure you have these Go libraries installed:  
```sh
go get fyne.io/fyne/v2
go get github.com/xuri/excelize/v2
```

---

## **Troubleshooting:**  
- Ensure your file is in the **Downloads folder** if entering the name manually.  
- Verify the file has an `.xlsx` extension.  
- Check error messages displayed in the dialog box.  

---
