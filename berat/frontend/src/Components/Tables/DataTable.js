import React, { useState } from 'react'
import { Table, Button, Modal, ModalHeader, ModalBody } from 'reactstrap';
import ModalForm from '../Modals/Modal'

function DataTable (props) {
  const [tempWeight, setTempWeight] = useState({})
  const [modalDetail, setModalDetail] = useState(false)
  const toggleDetail = (weight) => { 
    setTempWeight(weight)
    setModalDetail(!modalDetail) 
  }
  const closeBtnDetail = <button className="close" onClick={toggleDetail}>&times;</button>

  const weights = props.weights.map((weight, i) => {
    const isLast = i !== props.weights.length - 1

    return (
      <tr key={weight.date}>
        {isLast ? ( 
          <td style={{cursor:'pointer', color: '#3391ff'}} onClick={() => toggleDetail(weight)}>{weight.date}</td>
        ) : ( 
          <td>{weight.date}</td>
        )}
        <td>{weight.max}</td>
        <td>{weight.min}</td>
        <td>{weight.different}</td>
        <td>
          {isLast ? (
            <div style={{width:"150px"}}>
              <ModalForm 
                buttonLabel="Edit" 
                weight={weight} 
                getWeight={props.getWeight}
                createWeight={props.createWeight}
                updateWeight={props.updateWeight}
                deleteWeight={props.deleteWeight}
              />
              {' '}
              <Button 
                color="danger" 
                onClick={() => props.deleteWeight(weight.date)}>Delete</Button>
            </div>
          ) : null}
        </td>
      </tr>
    )
  })

  return (
    <div>
      <Table responsive hover>
        <thead>
          <tr>
            <th>Tanggal</th>
            <th>Max</th>
            <th>Min</th>
            <th>Perbedaan</th>
            <th>Pengaturan</th>
          </tr>
        </thead>
        <tbody>
          {weights}
        </tbody>
      </Table>
      <Modal isOpen={modalDetail} toggleDetail={toggleDetail}>
        <ModalHeader toggleDetail={toggleDetail} close={closeBtnDetail}>Detail {tempWeight.date}</ModalHeader>
        <ModalBody>
          <Table responsive hover>
          <thead>
            <tr>
              <th>Tanggal</th>
              <th>{tempWeight.date}</th>
            </tr>
          </thead>
          <tbody>
            <tr key="max">
              <td>Max</td>
              <td>{tempWeight.max}</td>
            </tr>
            <tr key="min">
              <td>Min</td>
              <td>{tempWeight.min}</td>
            </tr>
            <tr key="Perbedaan">
              <td>Perbedaan</td>
              <td>{tempWeight.different}</td>
            </tr>
          </tbody>
          </Table>
        </ModalBody>
      </Modal>
    </div> 
  )
}

export default DataTable