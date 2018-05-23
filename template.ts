/* eslint-disable */
// tslint:disable
const Icons = {
  {{ range . }}
  "{{ .Name }}": '{{ .Value }}',
  {{ end }}
};

export default Icons;