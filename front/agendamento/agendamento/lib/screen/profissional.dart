import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';

class ProfissionalScreen extends StatefulWidget {
  const ProfissionalScreen({super.key});

  @override
  State<ProfissionalScreen> createState() => _ProfissionalScreenState();
}

class _ProfissionalScreenState extends State<ProfissionalScreen> {
  @override
  Widget build(BuildContext context) {
    return SfCalendar(
          view: CalendarView.workWeek,
          dataSource: _getCalendarDataSource(),
          );
  }
}



class DataSource extends CalendarDataSource {
 DataSource(List<Appointment> source) {
   appointments = source;
 }
}


DataSource _getCalendarDataSource() {
   List<Appointment> appointments = <Appointment>[];
   appointments.add(
       Appointment(
         startTime: DateTime.now(),
         endTime: DateTime.now().add(
             const Duration(hours: 2)),
         isAllDay: false,
         subject: 'Meeting',
         color: Colors.blue,
         startTimeZone: 'America/Sao_Paulo',
         endTimeZone: 'America/Sao_Paulo'
       ));
  return DataSource(appointments);
}