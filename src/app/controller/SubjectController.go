package controller

import (
	"app/dataModel"
	"app/model"
	"app/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SubjectController struct {
	DataModel dataModel.ISubjectDataModel
}

func (this SubjectController) ReadAll(w http.ResponseWriter, r *http.Request) {
	result := this.DataModel.ReadAll()

	subject_json, err := json.Marshal(result)
	util.CheckErr(err)

	fmt.Fprintf(w, string(subject_json))
}

func (this SubjectController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectIdStr := vars["subjectId"]

	subjectId, err := strconv.ParseInt(subjectIdStr, 10, 64)
	util.CheckErr(err)

	result := this.DataModel.Read(subjectId)

	subject_json, err := json.Marshal(result)
	util.CheckErr(err)

	fmt.Fprintf(w, string(subject_json))
}

func (this SubjectController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectIdStr := vars["subjectId"]

	var temp struct {
		Name string `json:"url"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &temp)

	subjectId, err := strconv.ParseInt(subjectIdStr, 10, 64)
	util.CheckErr(err)

	subject := this.DataModel.Read(subjectId)
	subject.Name = temp.Name

	this.DataModel.Update(subject)
}

func (this SubjectController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectIdStr := vars["subjectId"]

	subjectId, err := strconv.ParseInt(subjectIdStr, 10, 64)
	util.CheckErr(err)

	subject := this.DataModel.Read(subjectId)

	this.DataModel.Delete(subject)
}

func (this SubjectController) Write(w http.ResponseWriter, r *http.Request) {
	var subject model.Subject

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &subject)

	this.DataModel.Write(subject)

	result := this.DataModel.ReadAll()

	subject_json, err := json.Marshal(result)
	util.CheckErr(err)

	fmt.Fprintf(w, string(subject_json))
}
