package main

type Attackable interface {
	Attack(BeAttackable)
}
type Player struct {
}
type Monster struct {
}

func (p *Player) Attack(target *BeAttackable) {

}
func (m *Monster) Attack(target *BeAttackable) {

}

type BeAttackable interface {
	Attackable()
}

func Attack(attacker *Attackable, defender *BeAttackable) {
	attacker.Attack(defender)
}

// type Report struct {
// }

// func (f *FinanceReport) MakeReport() *Report {

// }

// //확장성을 위한 틀 잡기
// type ReportSender interface {
// 	SendReport(*Report)
// }

// //세부사항
// type EmailReportSender struct {
// }

// func (f *EmailReportSender) SendReport(r *Report) {

// }

// type FileReportSender struct {
// }

// type HttpReportSender struct{

// }

// func(s *HttpReportSender) SendReport(r *Report){

// }
