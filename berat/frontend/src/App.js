import React, { useState, useEffect } from 'react';
import { Container, Row, Col } from 'reactstrap'
import DataTable from './Components/Tables/DataTable'
import ModalForm from './Components/Modals/Modal'

function App() {
  const [weights, setWeights] = useState([]);
  const [weight, setWeight] = useState({})

  const baseURL = "http://localhost:8000"

  const getWeights = () => {
    fetch(baseURL + '/weights')
      .then(response => {
        if (response.status >= 400 && response.status <= 600) {
          throw "failed get weights";
        } 

        return response.json()
      })
      .then(data => setWeights(data.Data))
      .catch(err => console.log(err))
  }

  const getWeight = (date) => {
    fetch(baseURL + '/weight?date=' + date)
      .then(response => {
        if (response.status >= 400 && response.status <= 600) {
          throw "failed get weight";
        }  
        
        return response.json()
      })
      .then(data => setWeight(data.Data))
      .catch(err => console.log(err))
  }

  const createWeight = (e, toggleForm, toggleError, setTextError) => {
    e.preventDefault()
    e.preventDefault()
    fetch(baseURL + '/weight', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        date: e.target.date.value,
        max: parseFloat(e.target.max.value),
        min: parseFloat(e.target.min.value)
      })
    })
      .then(response => {
        if (response.status >= 400 && response.status <= 600) {
          throw "failed create weight"
        } 
        
        return response.json()
      })
      .then(data => {
        setWeights(data.Data)
        toggleForm()
      })
      .catch(err => {
        setTextError(err)
        toggleForm()
        toggleError()
      })
  }

  const updateWeight = (e, toggleForm, toggleError, setTextError) => {
    e.preventDefault()
    fetch(baseURL + '/weight?date=' + e.target.date.value, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        max: parseFloat(e.target.max.value),
        min: parseFloat(e.target.min.value)
      })
    })
      .then(response => {
        if (response.status >= 400 && response.status <= 600) {
          throw "failed update weight"
        } 
        
        return response.json()
      })
      .then(data => {
        setWeights(data.Data)
        toggleForm()
      })
      .catch(err => {
        setTextError(err)
        toggleForm()
        toggleError()
      })
  }

  const deleteWeight = (date) => {
    fetch(baseURL + '/weight?date=' + date , {
      method: 'DELETE'
    })
      .then(response => {
        if (response.status >= 400 && response.status <= 600) {
          throw "failed delete weight";
        } 
        
        return response.json()
      })
      .then(data => setWeights(data.Data))
      .catch(err => console.log(err))
  }

  useEffect(() => {
    getWeights()
  }, [])

  return (
    <Container className="App">
      <Row>
        <Col>
          <h1 style={{margin: "20px 0"}}>CRUD Weight</h1>
        </Col>
      </Row>
      <Row>
        <Col>
          <DataTable 
            weights={weights} 
            weight={weight}
            getWeight={getWeight}
            createWeight={createWeight}
            updateWeight={updateWeight}
            deleteWeight={deleteWeight}
          />
        </Col>
      </Row>
      <Row>
        <Col>
          <ModalForm 
            buttonLabel="Add Weight" 
            getWeight={getWeight}
            createWeight={createWeight}
            updateWeight={updateWeight}
            deleteWeight={deleteWeight}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default App