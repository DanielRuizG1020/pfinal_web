 
 const apiUrl = 'http://localhost:8080/empleados';


 function showError(message) {
    const errorContenedor = document.createElement('div');
    errorContenedor.classList.add('error');
    errorContenedor.textContent = message;

    document.body.appendChild(errorContenedor);

    setTimeout(() => {
        document.body.removeChild(errorContenedor);
    }, 10000);
}

 function mostrarEmpleados(empleadosData) {
    employeeList.innerHTML = '';

    empleadosData.forEach(empleado => {
        const listItem = document.createElement('li');
        listItem.textContent = `ID: ${empleado.id}, Nombre: ${empleado.nombre}, Edad: ${empleado.edad} años, Nacionalidad: ${empleado.nacionalidad}, Ciudad: ${empleado.ciudad}, Casado: ${empleado.casado}, Discapacidad: ${empleado.discapacidad}`;
        employeeList.appendChild(listItem);
    });

}
 
document.addEventListener('DOMContentLoaded', function () {
    // Configurar eventos para el botón de carga de empleados
    const loadEmployeesButton = document.getElementById('loadEmployeesButton');
    loadEmployeesButton.addEventListener('click', function () {
        getEmployees();
    });
    const searchIdButton = document.getElementById('SearchID');
    searchIdButton.addEventListener('click', function () {
        searchEmployee();
    });

    const deleteEmployeeButton = document.getElementById('deleteEmployeeButton');
    deleteEmployeeButton.addEventListener('click', function () {
    deleteEmployee();
}
);
const createEmployeeButton = document.getElementById('createEmployeeButton');
createEmployeeButton.addEventListener('click', function () {
    createEmployee();
});

const updateEmployeeButton = document.getElementById('updateEmployeeButton');
updateEmployeeButton.addEventListener('click', function () {
    cupdateEmployee();
});

    // Función para obtener la lista de empleados (GET)
    async function getEmployees() {
        try {
            const response = await axios.get(apiUrl);
            const empleadosData = response.data;

            if (response.status === 200) {
                mostrarEmpleados(empleadosData);
            } else {
                showError('Error al obtener la lista de empleados.');
            }
        } catch (error) {
            showError('Error al obtener la lista de empleados.');
        }
    }

    // Función para mostrar la lista de empleados
    
});


async function searchEmployee() {
    const searchId = document.getElementById('searchId').value;

    try {
        const response = await axios.get(`${apiUrl}/${searchId}`);

        if (response.status === 200) {
            // Empleado encontrado, mostrar en la lista
            const empleadoEncontrado = response.data;
            mostrarEmpleados([empleadoEncontrado]);
        } else {
        
        }
    } catch (error) {
        console.error('Error al buscar el empleado:', error);
        
    }
}


// Función para eliminar un empleado por ID
async function deleteEmployee() {
    const deleteId = document.getElementById('deleteId').value;

    try {
        const response = await axios.delete(`${apiUrl}/${deleteId}`);

        if (response.status === 204) {
            // Operación exitosa, actualizar la lista de empleados
            getEmployees();
            showError('Eliminado con exito');
        } else {
            showError('Error al eliminar el empleado. Código de estado: ' + response.status);
        }
    } catch (error) {
        console.error('Error al eliminar el empleado:', error);
        showError('Error al eliminar el empleado. Consulta la consola para más detalles.');
    }
}

async function createEmployee() {
    const nombre = document.getElementById('nombre').value;
    const edad = parseInt(document.getElementById('edad').value, 10);
    const nacionalidad = document.getElementById('nacionalidad').value;
    const ciudad = document.getElementById('ciudad').value;
    const casado = document.getElementById('casado').checked;
    const discapacidad = document.getElementById('discapacidad').checked;

    // Agrega más campos según sea necesario

    try {
        const response = await axios.post(apiUrl, {
            nombre: nombre,
            edad: edad,
            nacionalidad: nacionalidad,
            ciudad: ciudad,
            casado: casado,
            discapacidad: discapacidad
        });

        if (response.status === 201) {
            // Operación exitosa, actualizar la lista de empleados
            getEmployees();
        } else {
            showError('Error al crear el empleado. Código de estado: ' + response.status);
        }
    } catch (error) {
        showError('Error al crear el empleado. Consulta la consola para más detalles.');
    }
}


async function updateEmployee() {
    const updateId = document.getElementById('updateId').value;
    const updateNombre = document.getElementById('updateNombre').value;
    const updateEdad = document.getElementById('updateEdad').value;
    const updateNacionalidad = document.getElementById('updateNacionalidad').value;
    const updateCiudad = document.getElementById('updateCiudad').value;
    const updateCasado = document.getElementById('updateCasado').checked;
    const updateDiscapacidad = document.getElementById('updateDiscapacidad').checked;

    // Obtener datos actuales del empleado
    const currentEmployee = await getEmployeeById(updateId);

    // Comparar y actualizar solo los campos que han cambiado
    const updatedData = {
        nombre: updateNombre || currentEmployee.nombre,
        edad: updateEdad || currentEmployee.edad,
        nacionalidad: updateNacionalidad || currentEmployee.nacionalidad,
        ciudad: updateCiudad || currentEmployee.ciudad,
        casado: updateCasado,
        discapacidad: updateDiscapacidad
    };

    try {
        const response = await axios.patch(`${apiUrl}/${updateId}`, updatedData);

        if (response.status === 200) {
            // Operación exitosa, actualizar la lista de empleados
            getEmployees();
        } else {
            showError('Error al actualizar el empleado. Código de estado: ' + response.status);
        }
    } catch (error) {
        console.error('Error al actualizar el empleado:', error);
        showError('Error al actualizar el empleado. Consulta la consola para más detalles.');
    }
}

// Función para obtener datos actuales de un empleado por ID
async function getEmployeeById(employeeId) {
    try {
        const response = await axios.get(`${apiUrl}/${employeeId}`);
        return response.data;
    } catch (error) {
        console.error('Error al obtener datos del empleado por ID:', error);
    }
}