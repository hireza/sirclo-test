import React from 'react';
import { Button, Form, FormGroup, Label, Input } from 'reactstrap';

function AddEditForm (props) {
  return (
    <Form onSubmit={props.weight ? (e) => props.updateWeight(e, props.toggleForm, props.toggleError, props.setTextError) : (e) => props.createWeight(e, props.toggleForm, props.toggleError, props.setTextError)}>
      <FormGroup>
        <Label for="date">Tanggal</Label>
        <Input type="text" name="date" id="date" placeholder="DD-MM-YYYY"  defaultValue={props.weight ? props.weight.date : ""} />
      </FormGroup>
      <FormGroup>
        <Label for="max">Max</Label>
        <Input type="number" name="max" id="max" defaultValue={props.weight ? props.weight.max : ""} />
      </FormGroup>
      <FormGroup>
        <Label for="min">Min</Label>
        <Input type="number" name="min" id="min" defaultValue={props.weight ? props.weight.min : ""} />
      </FormGroup>
      <Button>Submit</Button>
    </Form>
  );
}

export default AddEditForm