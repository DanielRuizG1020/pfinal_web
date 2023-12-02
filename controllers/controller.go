package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"proyectofinal/models"
	repositorio "proyectofinal/repository"
)

var (
	updateQuery = "UPDATE empleados SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM empleados WHERE id=$1;"
	selectQuery = "SELECT id, nombre, edad, nacionalidad, ciudad, casado, discapacidad FROM empleados WHERE id=$1;"
	listQuery   = "SELECT id, nombre, edad, nacionalidad, ciudad, casado, discapacidad FROM empleados limit $1 offset $2"
	createQuery = "INSERT INTO empleados (nombre, edad, nacionalidad, ciudad, casado, discapacidad) VALUES (:nombre, :edad, :nacionalidad, :ciudad, :casado, :discapacidad) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Empleado]
}

func NewController(repo repositorio.Repository[models.Empleado]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnEmpleado(reqBody []byte, id string) error {
	nuevosValoresEmpleado := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresEmpleado)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un empleado, con error: %s", err.Error())
	}

	if len(nuevosValoresEmpleado) == 0 {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un empleado, con error: %s", err.Error())
	}

	query := construirUpdateQuery(nuevosValoresEmpleado)
	nuevosValoresEmpleado["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresEmpleado)
	if err != nil {
		log.Printf("fallo al actualizar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un empleado, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columns := []string{}
	for key := range nuevosValores {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarUnEmpleado(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un empleado, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un empleado, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnEmpleado(id string) ([]byte, error) {
	empleado, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un empleado, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un empleado, con error: %s", err.Error())
	}

	amigoJson, err := json.Marshal(empleado)
	if err != nil {
		log.Printf("fallo al leer un empleado, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un empleado, con error: %s", err.Error())
	}
	return amigoJson, nil
}

func (c *Controller) LeerEmpleados(limit, offset int) ([]byte, error) {
	empleados, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer empleados, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer empleados, con error: %s", err.Error())
	}

	jsonEmpleados, err := json.Marshal(empleados)
	if err != nil {
		log.Printf("fallo al leer empleados, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer empleados, con error: %s", err.Error())
	}
	return jsonEmpleados, nil
}

func (c *Controller) CrearEmpleado(reqBody []byte) (int64, error) {
	nuevoEmpleado := &models.Empleado{}
	err := json.Unmarshal(reqBody, nuevoEmpleado)
	if err != nil {
		log.Printf("fallo al crear un nuevo empleado, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo empleado, con error: %s", err.Error())
	}

	valoresColumnasNuevoAmigo := map[string]any{
		"nombre":       nuevoEmpleado.Nombre,
		"edad":         nuevoEmpleado.Edad,
		"nacionalidad": nuevoEmpleado.Nacionalidad,
		"ciudad":       nuevoEmpleado.Ciudad,
		"casado":       nuevoEmpleado.Casado,
		"discapacidad": nuevoEmpleado.Discapacidad,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoAmigo)
	if err != nil {
		log.Printf("fallo al crear un nuevo empleado, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo empleado, con error: %s", err.Error())
	}
	return nuevoId, nil
}
