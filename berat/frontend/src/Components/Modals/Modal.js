import React, { useState } from 'react'
import { Button, Modal, ModalHeader, ModalBody } from 'reactstrap'
import AddEditForm from '../Forms/AddEditForm'

function ModalForm (props) {
  const [modalForm, setModalForm] = useState(false)
  const [modalError, setModalError] = useState(false)
  const [textError, setTextError] = useState("")

  const toggleForm = () => { setModalForm(!modalForm) }
  const toggleError = () => { setModalError(!modalError) }

  const closeBtn = <button className="close" onClick={toggleForm}>&times;</button>
  const closeBtnError = <button className="close" onClick={toggleError}>&times;</button>

  const label = props.buttonLabel

  let button = ''
  let title = ''

  if(label === 'Edit'){
    button = <Button
              color="warning"
              onClick={toggleForm}
              style={{float: "left", marginRight:"10px"}}>{label}</Button>
    title = 'Edit Weight'
  } else {
    button = <Button
              color="success"
              onClick={toggleForm}
              style={{float: "left", marginRight:"10px"}}>{label}</Button>
    title = 'Add New Weight'
  }

  return (
    <div>
      {button}
      <Modal isOpen={modalForm} toggleForm={toggleForm}>
        <ModalHeader toggleForm={toggleForm} close={closeBtn}>{title}</ModalHeader>
        <ModalBody>
          <AddEditForm
            weight={props.weight} 
            createWeight={props.createWeight}
            updateWeight={props.updateWeight}
            toggleForm={toggleForm}
            toggleError={toggleError}
            setTextError={setTextError}
          />
        </ModalBody>
      </Modal>
      <Modal isOpen={modalError} toggleError={toggleError}>
        <ModalHeader toggleError={toggleError} close={closeBtnError}>{title}</ModalHeader>
        <ModalBody>
          {textError}
        </ModalBody>
      </Modal>
    </div>
  )
}

export default ModalForm